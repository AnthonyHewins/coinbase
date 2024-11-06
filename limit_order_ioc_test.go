package coinbase

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestLimitOrderIOCJson(t *testing.T) {
	marshalTest(t, &LimitOrderIOC{
		BaseSize:   decimal.NewFromFloat(19203),
		LimitPrice: decimal.NewFromFloat(43254),
	})
}
