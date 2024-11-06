package coinbase

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestStopOrderGTDJSON(t *testing.T) {
	marshalTest(t, &StopLimitOrderGTD{
		BaseSize:   decimal.NewFromFloat(8210938.2398),
		LimitPrice: decimal.NewFromFloat(821389),
		Stop:       decimal.NewFromFloat(1231),
		EndTime:    time.Date(1, 1, 1, 1, 11, 1, 0, time.UTC),
		Side:       SideBuy,
	})
}
