package coinbase

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestTriggerBracketOrderGTCJson(t *testing.T) {
	marshalTest(t, &TriggerBracketOrderGTC{
		BaseSize:         decimal.NewFromFloat(32425.43),
		LimitPrice:       decimal.NewFromFloat(3254.54),
		StopTriggerPrice: decimal.NewFromFloat(123.541),
	})
}
