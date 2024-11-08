package coinbase

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
)

type limitGtdWrapper struct {
	Data *orderData `json:"limit_limit_gtd"`
}

type LimitOrderGTD struct {
	// The amount of the first Asset in the Trading Pair. For example, on the
	// BTC-USD Order Book, BTC is the Base Asset.
	BaseSize decimal.Decimal

	// The specified price, or better, that the Order should be executed at. A
	// Buy Order will execute at or lower than the limit price. A Sell Order
	// will execute at or higher than the limit price.
	LimitPrice decimal.Decimal
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
		Data: &orderData{
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
