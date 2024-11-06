package coinbase

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestTriggerBracketOrderGTDJson(t *testing.T) {
	marshalTest(t, &TriggerBracketOrderGTD{
		BaseSize:         decimal.NewFromFloat(32425.43),
		LimitPrice:       decimal.NewFromFloat(3254.54),
		EndTime:          time.Date(1, 1, 2, 3, 124, 432, 0, time.UTC),
		StopTriggerPrice: decimal.NewFromFloat(123.541),
	})
}
