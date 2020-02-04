package types

import (
	"encoding/json"
	"strconv"
)

type PriceWrapper string

type RefreshType struct {
	Ask     PriceWrapper `json:"ask"`
	AskSize uint16       `json:"asksize"`
	Avg     PriceWrapper `json:"avg"`
	Bid     PriceWrapper `json:"bid"`
	BidSize uint16       `json:"bidsize"`
	Delta   string       `json:"delta"`
	High    PriceWrapper `json:"high"`
	Last    PriceWrapper `json:"last"`
	Low     PriceWrapper `json:"low"`
	Refresh uint16       `json:"refresh"`
	Stueck  uint16       `json:"stueck"`
	Umsatz  uint64       `json:"umsatz"`
}

func (w *PriceWrapper) UnmarshalJSON(data []byte) (err error) {
	if f, err := strconv.ParseFloat(string(data), 64); err == nil {
		str := strconv.FormatFloat(f, 'f', -1, 64)
		*w = PriceWrapper(str)
		return nil
	}

	var str string
	err = json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	*w = PriceWrapper(str)
	return nil
}
