package coinbase

import (
	"encoding/json"
)

type triggerBracketGTCWrapper struct {
	Data orderData `json:"trigger_bracket_gtc"`
}

type TriggerBracketOrderGTC struct {
	BaseSize   string `json:"base_size"`
	LimitPrice string `json:"limit_price"`

	// The price level (in quote currency) where the position will be exited.
	// When triggered, a stop limit order is automatically placed with a limit
	// price 5% higher for BUYS and 5% lower for SELLS.
	StopTriggerPrice string `json:"stop_trigger_price"`
}

func (t *TriggerBracketOrderGTC) OrderType() OrderType { return TriggerBracketGTC }

func (t *TriggerBracketOrderGTC) MarshalJSON() ([]byte, error) {
	return json.Marshal(triggerBracketGTCWrapper{
		Data: orderData{
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
