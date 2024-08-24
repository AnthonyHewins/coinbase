package coinbase

import (
	"testing"
	"time"
)

func TestTriggerBracketOrderGTDJson(t *testing.T) {
	marshalTest(t, &TriggerBracketOrderGTD{
		BaseSize:         "32425.43",
		LimitPrice:       "3254.54",
		EndTime:          time.Date(1, 1, 2, 3, 124, 432, 0, time.UTC),
		StopTriggerPrice: "123.541",
	})
}
