package coinbase

import "testing"

func TestMarketOrderJSON(t *testing.T) {
	marshalTest(t, &MarketOrder{
		QuoteSize: "1234",
		BaseSize:  "3423.5324",
	})
}
