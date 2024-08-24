package coinbase

import "testing"

func TestLimitOrderIOCJson(t *testing.T) {
	marshalTest(t, &LimitOrderIOC{
		BaseSize:   "19203",
		LimitPrice: "43254",
	})
}
