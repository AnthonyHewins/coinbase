package coinbase

import "testing"

func TestTriggerBracketOrderGTCJson(t *testing.T) {
	marshalTest(t, &TriggerBracketOrderGTC{
		BaseSize:         "32425.43",
		LimitPrice:       "3254.54",
		StopTriggerPrice: "123.541",
	})
}
