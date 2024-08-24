package coinbase

import (
	"encoding/json"
	"time"
)

type stopLimitGTDWrapper struct {
	Data orderData `json:"stop_limit_stop_limit_gtd"`
}

type StopLimitOrderGTD struct {
	BaseSize   string
	LimitPrice string
	Stop       string
	EndTime    time.Time

	// on which side should the limit trigger.
	// if you specify Buy then the stop is triggered once the stop EXCEEDS.
	// if you specify Sell then the stop is triggered once the stop FALLS BELOW
	// If you don't specify one, the payload to coinbase will not have one,
	// and will result in an error
	Side Side
}

func (s *StopLimitOrderGTD) OrderType() OrderType { return StopLimitGTD }

func (s *StopLimitOrderGTD) MarshalJSON() ([]byte, error) {
	var t *time.Time
	if !s.EndTime.IsZero() {
		t = &s.EndTime
	}

	return json.Marshal(stopLimitGTDWrapper{
		Data: orderData{
			BaseSize:   s.BaseSize,
			LimitPrice: s.LimitPrice,
			End:        t,
			Stop:       s.Stop,
			Side:       s.Side.toStopDirectionStr(),
		},
	})
}

func (s *StopLimitOrderGTD) UnmarshalJSON(buf []byte) error {
	var w stopLimitGTDWrapper
	if err := json.Unmarshal(buf, &w); err != nil {
		return err
	}

	side, err := w.Data.stopSide()
	if err != nil {
		return err
	}

	*s = StopLimitOrderGTD{
		BaseSize:   w.Data.BaseSize,
		LimitPrice: w.Data.LimitPrice,
		Stop:       w.Data.Stop,
		EndTime:    w.Data.end(),
		Side:       side,
	}

	return nil
}
