package coinbase

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestLimitOrderFOKJSON(t *testing.T) {
	marshalTest(t, &LimitOrderFOK{
		BaseSize:   decimal.NewFromFloat(1233),
		LimitPrice: decimal.NewFromFloat(5436.23),
	})
}
