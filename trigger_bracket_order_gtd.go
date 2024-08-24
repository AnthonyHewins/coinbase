package coinbase

import (
	"encoding/json"
	"time"
)

type triggerBracketGTDWrapper struct {
	Data orderData `json:"trigger_bracket_gtd"`
}

type TriggerBracketOrderGTD struct {
	BaseSize   string    `json:"base_size"`
	LimitPrice string    `json:"limit_price"`
	EndTime    time.Time `json:"end_time"`

	// The price level (in quote currency) where the position will be exited.
	// When triggered, a stop limit order is automatically placed with a limit
	// price 5% higher for BUYS and 5% lower for SELLS.
	StopTriggerPrice string `json:"stop_trigger_price"`
}

func (t *TriggerBracketOrderGTD) OrderType() OrderType { return TriggerBracketGTD }

func (t *TriggerBracketOrderGTD) MarshalJSON() ([]byte, error) {
	var s *time.Time
	if !t.EndTime.IsZero() {
		s = &t.EndTime
	}

	return json.Marshal(triggerBracketGTDWrapper{
		Data: orderData{
			BaseSize:         t.BaseSize,
			LimitPrice:       t.LimitPrice,
			End:              s,
			StopTriggerPrice: t.StopTriggerPrice,
		},
	})
}

func (t *TriggerBracketOrderGTD) UnmarshalJSON(b []byte) error {
	var x triggerBracketGTDWrapper
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}

	*t = TriggerBracketOrderGTD{
		BaseSize:         x.Data.BaseSize,
		LimitPrice:       x.Data.LimitPrice,
		EndTime:          x.Data.end(),
		StopTriggerPrice: x.Data.StopTriggerPrice,
	}

	return nil
}
