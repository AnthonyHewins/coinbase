package coinbase

import "testing"

func TestLimitOrderGTCJSON(t *testing.T) {
	marshalTest(t, &LimitOrderGTC{
		BaseSize:   "214532",
		LimitPrice: "4325454",
		PostOnly:   true,
	})
}
