package coinbase

import (
	"encoding/json"
	"time"
)

type limitGtdWrapper struct {
	Data orderData `json:"limit_limit_gtd"`
}

type LimitOrderGTD struct {
	BaseSize   string
	LimitPrice string
	EndTime    time.Time

	// Enable or disable Post-only Mode. When enabled, only Maker Orders will be
	// posted to the Order Book. Orders that will be posted as a Taker Order
	// will be rejected.
	PostOnly bool
}

func (l *LimitOrderGTD) OrderType() OrderType { return LimitGTD }

func (l *LimitOrderGTD) MarshalJSON() ([]byte, error) {
	var t *time.Time
	if !l.EndTime.IsZero() {
		t = &l.EndTime
	}

	return json.Marshal(limitGtdWrapper{
		Data: orderData{
			BaseSize:   l.BaseSize,
			LimitPrice: l.LimitPrice,
			PostOnly:   l.PostOnly,
			End:        t,
		},
	})
}

func (l *LimitOrderGTD) UnmarshalJSON(b []byte) error {
	var x limitGtdWrapper
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}

	*l = LimitOrderGTD{
		BaseSize:   x.Data.BaseSize,
		LimitPrice: x.Data.LimitPrice,
		PostOnly:   x.Data.PostOnly,
		EndTime:    x.Data.end(),
	}

	return nil
}
