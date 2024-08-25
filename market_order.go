package coinbase

import "encoding/json"

type marketOrderWrapper struct {
	Outer orderData `json:"market_market_ioc"`
}

type MarketOrder struct {
	// The amount of the second Asset in the Trading Pair. For example, on the
	// BTC/USD Order Book, USD is the Quote Asset.
	QuoteSize string

	// The amount of the first Asset in the Trading Pair. For example, on the
	// BTC-USD Order Book, BTC is the Base Asset.
	BaseSize string
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
