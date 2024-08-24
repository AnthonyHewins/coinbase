package coinbase

import (
	"testing"
	"time"
)

func TestLimitOrderGTDJSON(t *testing.T) {
	marshalTest(t, &LimitOrderGTD{
		BaseSize:   "213423.54",
		LimitPrice: "325343",
		EndTime:    time.Date(3, 3, 4, 5, 3, 3, 0, time.UTC),
		PostOnly:   true,
	})
}
