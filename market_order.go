package coinbase

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type marketOrderWrapper struct {
	Data *orderData `json:"market_market_ioc"`
}

type MarketOrder struct {
	// The amount of the second Asset in the Trading Pair. For example, on the
	// BTC/USD Order Book, USD is the Quote Asset.
	QuoteSize decimal.Decimal

	// The amount of the first Asset in the Trading Pair. For example, on the
	// BTC-USD Order Book, BTC is the Base Asset.
	BaseSize decimal.Decimal
}

func (m *MarketOrder) OrderType() OrderType { return Market }

func (m *MarketOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(marketOrderWrapper{
		Data: &orderData{
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
		QuoteSize: x.Data.QuoteSize,
		BaseSize:  x.Data.BaseSize,
	}

	return nil
}
