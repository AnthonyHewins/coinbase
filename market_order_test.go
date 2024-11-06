package coinbase

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestMarketOrderJSON(t *testing.T) {
	marshalTest(t, &MarketOrder{
		QuoteSize: decimal.NewFromFloat(1234),
		BaseSize:  decimal.NewFromFloat(3423.5324),
	})
}
