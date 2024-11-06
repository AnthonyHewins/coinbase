package coinbase

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type triggerBracketGTCWrapper struct {
	Data *orderData `json:"trigger_bracket_gtc"`
}

type TriggerBracketOrderGTC struct {
	// The amount of the first Asset in the Trading Pair. For example, on the
	// BTC-USD Order Book, BTC is the Base Asset.
	BaseSize decimal.Decimal `json:"base_size"`

	// The specified price, or better, that the Order should be executed at. A
	// Buy Order will execute at or lower than the limit price. A Sell Order
	// will execute at or higher than the limit price.
	LimitPrice decimal.Decimal `json:"limit_price"`

	// The price level (in quote currency) where the position will be exited.
	// When triggered, a stop limit order is automatically placed with a limit
	// price 5% higher for BUYS and 5% lower for SELLS.
	StopTriggerPrice decimal.Decimal `json:"stop_trigger_price"`
}

func (t *TriggerBracketOrderGTC) OrderType() OrderType { return TriggerBracketGTC }

func (t *TriggerBracketOrderGTC) MarshalJSON() ([]byte, error) {
	return json.Marshal(triggerBracketGTCWrapper{
		Data: &orderData{
			BaseSize:         t.BaseSize,
			LimitPrice:       t.LimitPrice,
			StopTriggerPrice: t.StopTriggerPrice,
		},
	})
}

func (t *TriggerBracketOrderGTC) UnmarshalJSON(b []byte) error {
	var x triggerBracketGTCWrapper
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}

	*t = TriggerBracketOrderGTC{
		BaseSize:         x.Data.BaseSize,
		LimitPrice:       x.Data.LimitPrice,
		StopTriggerPrice: x.Data.StopTriggerPrice,
	}

	return nil
}
