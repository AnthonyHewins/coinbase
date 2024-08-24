package coinbase

import (
	"testing"
)

func TestStopOrderGTCJSON(t *testing.T) {
	marshalTest(t, &StopLimitOrderGTC{
		BaseSize:   "8210938.2398",
		LimitPrice: "821389",
		Stop:       "1231",
		Side:       SideBuy,
	})
}
