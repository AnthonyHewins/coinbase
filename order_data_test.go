package coinbase

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
)

func marshalTest[X OrderConfig](t *testing.T, x X) {
	tt := assert.New(t)

	buf, err := json.Marshal(x)
	if !tt.NoError(err, "should not fail marshal") {
		return
	}

	snaps.MatchJSON(t, buf)

	var y X
	if err = json.Unmarshal(buf, &y); !tt.NoError(err, "should not fail unmarshal") {
		return
	}

	tt.Equal(x, y, "marshal and unmarshal must be bidi")
}

func TestDiscoverConfig(mainTest *testing.T) {
	end := time.Date(23, 12, 30, 1, 1, 1, 0, time.UTC)
	testCases := []struct {
		orderType string
		c         OrderConfig
	}{
		{
			orderType: "MARKET",
			c:         &MarketOrder{QuoteSize: "123", BaseSize: "123"},
		},
		{
			orderType: "LIMIT",
			c:         &LimitOrderIOC{BaseSize: "1234", LimitPrice: "1234"},
		},
		{
			orderType: "LIMIT",
			c:         &LimitOrderFOK{BaseSize: "1234", LimitPrice: "1234"},
		},
		{
			orderType: "LIMIT",
			c: &LimitOrderGTC{
				BaseSize:   "1234",
				LimitPrice: "1234",
				PostOnly:   true,
			},
		},
		{
			orderType: "LIMIT",
			c: &LimitOrderGTD{
				BaseSize:   "1234",
				LimitPrice: "1234",
				EndTime:    end,
				PostOnly:   true,
			},
		},
		{
			orderType: "STOP",
			c: &StopLimitOrderGTC{
				BaseSize:   "82103",
				LimitPrice: "34254",
				Stop:       "324",
				Side:       SideBuy,
			},
		},
		{
			orderType: "STOP",
			c: &StopLimitOrderGTD{
				BaseSize:   "82103",
				LimitPrice: "34254",
				Stop:       "324",
				EndTime:    end,
				Side:       SideBuy,
			},
		},
		{
			orderType: "BRACKET",
			c: &TriggerBracketOrderGTC{
				BaseSize:         "2343",
				LimitPrice:       "5544",
				StopTriggerPrice: "123",
			},
		},
		{
			orderType: "BRACKET",
			c: &TriggerBracketOrderGTD{
				BaseSize:         "2343",
				LimitPrice:       "5544",
				EndTime:          end,
				StopTriggerPrice: "123",
			},
		},
	}

	for _, tc := range testCases {
		mainTest.Run(fmt.Sprintf("discovers type %s", tc.c.OrderType()), func(tt *testing.T) {
			t := assert.New(tt)

			buf, err := json.Marshal(tc.c)
			if err != nil {
				t.Fail("failed marshal to begin test: %s", err)
			}

			actual, actualErr := discoverConfig(tc.orderType, buf)

			if t.NoError(actualErr, "should not error when a value is expected") {
				t.Equal(tc.c, actual)
			}
		})
	}
}