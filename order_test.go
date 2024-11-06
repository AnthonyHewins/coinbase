package coinbase

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func mustParse(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}

	return t
}

func TestUnmarshalOrder(mainTest *testing.T) {
	testCases := []struct {
		name        string
		arg         string
		expected    Order
		expectedErr string
	}{
		{
			name: "random example",
			arg: `{
			"order_id":"4c22e9e1-54cc-449d-97da-8f6ca6782e3e",
			"product_id":"BTC-USD",
			"user_id":"e0b87645-7245-5705-8f29-6550318cda60",
			"order_configuration":{"limit_limit_gtd":{"base_size":"0.00000001", "limit_price":"0.01", "end_time":"2024-08-30T13:07:40.823884203Z", "post_only":false}},
			"side":"BUY",
			"client_order_id":"113b22a6-1fb3-4187-9d0d-a4476e91913c",
			"status":"OPEN",
			"time_in_force":"GOOD_UNTIL_DATE_TIME",
			"created_time":"2024-08-30T13:07:25.942018Z",
			"completion_percentage":"0",
			"filled_size":"0",
			"average_filled_price":"0",
			"fee":"",
			"number_of_fills":"0",
			"filled_value":"0",
			"pending_cancel":false,
			"size_in_quote":false,
			"total_fees":"0",
			"size_inclusive_of_fees":false,
			"total_value_after_fees":"0",
			"trigger_status":"INVALID_ORDER_TYPE",
			"order_type":"LIMIT",
			"reject_reason":"REJECT_REASON_UNSPECIFIED",
			"settled":false,
			"product_type":"SPOT",
			"reject_message":"",
			"cancel_message":"",
			"order_placement_source":"RETAIL_ADVANCED",
			"outstanding_hold_amount":"0.00000000010075",
			"is_liquidation":false,
			"last_fill_time":null,
			"edit_history":[],
			"leverage":"",
			"margin_type":"UNKNOWN_MARGIN_TYPE",
			"retail_portfolio_id":"e0b87645-7245-5705-8f29-6550318cda60",
			"originating_order_id":"",
			"attached_order_id":""}`,
			expected: Order{
				ID:        uuid.MustParse("4c22e9e1-54cc-449d-97da-8f6ca6782e3e"),
				ProductID: "BTC-USD",
				User:      "e0b87645-7245-5705-8f29-6550318cda60",
				Config: &LimitOrderGTD{
					BaseSize:   decimal.NewFromFloat(0.00000001),
					LimitPrice: decimal.NewFromFloat(0.01),
					EndTime:    mustParse("2024-08-30T13:07:40.823884203Z"),
				},
				Side:                  SideBuy,
				IdemKey:               "113b22a6-1fb3-4187-9d0d-a4476e91913c",
				Status:                StatusOpen,
				TIF:                   TIFGoodUntilDateTime,
				Created:               mustParse("2024-08-30T13:07:25.942018Z"),
				TriggerStatus:         TriggerStatusInvalidOrderType,
				ProductType:           ProductTypeSpot,
				Edits:                 []EditHistory{},
				OrderPlacementSrc:     OrderPlacementSrcRetailAdvanced,
				OutstandingHoldAmount: "0.00000000010075",
				RetailPortfolioID:     "e0b87645-7245-5705-8f29-6550318cda60",
				Completion:            "0",
				FilledSize:            "0",
				AvgFillPrice:          "0",
				NumberOfFills:         "0",
				FilledValue:           "0",
				TotalFees:             "0",
				TotalValueAfterFees:   "0",
			},
			expectedErr: "",
		},
	}

	for _, tc := range testCases {
		mainTest.Run(tc.name, func(tt *testing.T) {
			t := assert.New(tt)
			var o Order
			actualErr := json.Unmarshal([]byte(tc.arg), &o)

			if tc.expectedErr != "" {
				if t.NoError(actualErr, "should not return happy path value when error is expected") {
					t.EqualError(actualErr, tc.expectedErr)
				}
				return
			}

			if t.NoError(actualErr, "should not error when a value is expected") {
				t.Equal(tc.expected, o, tc.name)
			}
		})
	}
}
