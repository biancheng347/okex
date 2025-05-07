package main

import (
	"context"
	"encoding/json"
	"github.com/biancheng347/okex"
	"github.com/biancheng347/okex/api"
	"github.com/biancheng347/okex/events"
	"github.com/biancheng347/okex/events/private"
	ws_private_requests "github.com/biancheng347/okex/requests/ws/private"
	"log"
)

func main() {
	apiKey := ""
	secretKey := ""
	passphrase := ""
	dest := okex.NormalServer // The main API server
	ctx := context.Background()
	client, err := api.NewClient(ctx, apiKey, secretKey, passphrase, dest)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Starting")
	errChan := make(chan *events.Error)
	subChan := make(chan *events.Subscribe)
	uSubChan := make(chan *events.Unsubscribe)
	logChan := make(chan *events.Login)
	sucChan := make(chan *events.Success)
	client.BWs.SetChannels(errChan, subChan, uSubChan, logChan, sucChan)

	//cCh := make(chan *public.Candlesticks)
	//err = client.BWs.Public.Candlesticks(ws_public_requests.Candlesticks{
	//	InstID:  "BTC-USDT-SWAP",
	//	Channel: okex.CandleStick5m,
	//}, cCh)

	//err = client.BWs.Public.Candlesticks(ws_public_requests.Candlesticks{
	//	InstID:  "ETH-USDT-SWAP",
	//	Channel: okex.CandleStick5m,
	//}, cCh)

	pCh := make(chan *private.Position)

	data := map[string]string{
		"updateInterval": "2000",
	}
	jsonData, _ := json.MarshalIndent(data, "", "    ") // 使用缩进格式化
	extra := string(jsonData)

	err = client.BWs.Private.PositionExtra(ws_private_requests.PositionExtra{
		InstID:      "BTC-USDT-SWAP",
		InstType:    "SWAP",
		ExtraParams: extra,
	}, pCh)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		select {
		case <-logChan:
			log.Print("[Authorized]")
		case success := <-sucChan:
			log.Printf("[SUCCESS]\t%+v", success)
		case sub := <-subChan:
			channel, _ := sub.Arg.Get("channel")
			log.Printf("[Subscribed]\t%s", channel)
		case uSub := <-uSubChan:
			channel, _ := uSub.Arg.Get("channel")
			log.Printf("[Unsubscribed]\t%s", channel)
		case err := <-client.Ws.ErrChan:
			log.Printf("[Error]\t%+v", err)
			for _, datum := range err.Data {
				log.Printf("[Error]\t\t%+v", datum)
			}
		case p := <-pCh:
			for _, position := range p.Positions {
				log.Printf("[Position]\t%+v", position)
			}
		//case c := <-cCh:
		//	instId, ok := c.Arg.Get("instId")
		//	if ok {
		//		for _, candle := range c.Candles {
		//			log.Printf("[Candlesticks: %s]\t%+v", instId, candle)
		//		}
		//	}
		case b := <-client.Ws.DoneChan:
			log.Printf("[End]:\t%v", b)
			return
		}
	}
}
