package coinbase

import "testing"

func TestLimitOrderFOKJSON(t *testing.T) {
	marshalTest(t, &LimitOrderFOK{
		BaseSize:   "1233",
		LimitPrice: "5436.23",
	})
}
