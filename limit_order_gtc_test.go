package coinbase

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestLimitOrderGTCJSON(t *testing.T) {
	marshalTest(t, &LimitOrderGTC{
		BaseSize:   decimal.NewFromFloat(214532),
		LimitPrice: decimal.NewFromFloat(4325454),
		PostOnly:   true,
	})
}
