package private

import "github.com/biancheng347/okex"

type (
	Account struct {
		Ccy string `json:"ccy,omitempty"`
	}
	Position struct {
		Uly      string              `json:"uly,omitempty"`
		InstID   string              `json:"instId,omitempty"`
		InstType okex.InstrumentType `json:"instType"`
	}
	PositionExtra struct {
		Uly         string              `json:"uly,omitempty"`
		InstID      string              `json:"instId,omitempty"`
		InstType    okex.InstrumentType `json:"instType"`
		ExtraParams string              `json:"extraParams,omitempty"`
	}
	Order struct {
		Uly      string              `json:"uly,omitempty"`
		InstID   string              `json:"instId,omitempty"`
		InstType okex.InstrumentType `json:"instType"`
	}
	AlgoOrder struct {
		Uly      string              `json:"uly,omitempty"`
		InstID   string              `json:"instId,omitempty"`
		InstType okex.InstrumentType `json:"instType"`
	}
)
