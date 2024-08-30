package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBidAsk(mainTest *testing.T) {
	c := testClient()

	t := assert.New(mainTest)

	pairs, err := c.BidAsk(context.Background(), "BTC-USD")
	if !t.NoError(err, "should not fail request to get bid/ask") {
		return
	}

	if !t.NotEmpty(pairs, "should return BTC-USD when requested") {
		return
	}
}
