package coinbase

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestStopOrderGTCJSON(t *testing.T) {
	marshalTest(t, &StopLimitOrderGTC{
		BaseSize:   decimal.NewFromFloat(8210938.2398),
		LimitPrice: decimal.NewFromFloat(821389),
		Stop:       decimal.NewFromFloat(1231),
		Side:       SideBuy,
	})
}
