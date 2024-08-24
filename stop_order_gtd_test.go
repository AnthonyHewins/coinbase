package coinbase

import (
	"testing"
	"time"
)

func TestStopOrderGTDJSON(t *testing.T) {
	marshalTest(t, &StopLimitOrderGTD{
		BaseSize:   "8210938.2398",
		LimitPrice: "821389",
		Stop:       "1231",
		EndTime:    time.Date(1, 1, 1, 1, 11, 1, 0, time.UTC),
		Side:       SideBuy,
	})
}
