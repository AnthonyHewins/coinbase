package coinbase

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestLimitOrderGTDJSON(t *testing.T) {
	marshalTest(t, &LimitOrderGTD{
		BaseSize:   decimal.NewFromFloat(213423.54),
		LimitPrice: decimal.NewFromFloat(325343),
		EndTime:    time.Date(3, 3, 4, 5, 3, 3, 0, time.UTC),
		PostOnly:   true,
	})
}
