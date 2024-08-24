package coinbase

import "encoding/json"

type marketOrderWrapper struct {
	Outer orderData `json:"market_market_ioc"`
}

type MarketOrder struct {
	QuoteSize string
	BaseSize  string
}

func (m *MarketOrder) OrderType() OrderType { return Market }

func (m *MarketOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(marketOrderWrapper{
		Outer: orderData{
			QuoteSize: m.QuoteSize,
			BaseSize:  m.BaseSize,
		},
	})
}

func (m *MarketOrder) UnmarshalJSON(b []byte) error {
	var x marketOrderWrapper
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}

	*m = MarketOrder{
		QuoteSize: x.Outer.QuoteSize,
		BaseSize:  x.Outer.BaseSize,
	}

	return nil
}
