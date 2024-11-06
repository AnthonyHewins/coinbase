package coinbase

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
)

type stopLimitGTDWrapper struct {
	Data *orderData `json:"stop_limit_stop_limit_gtd"`
}

type StopLimitOrderGTD struct {
	// The amount of the first Asset in the Trading Pair. For example, on the
	// BTC-USD Order Book, BTC is the Base Asset.
	BaseSize decimal.Decimal

	// The specified price, or better, that the Order should be executed at. A
	// Buy Order will execute at or lower than the limit price. A Sell Order
	// will execute at or higher than the limit price.
	LimitPrice decimal.Decimal
	Stop       decimal.Decimal
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
		Data: &orderData{
			BaseSize:      s.BaseSize,
			LimitPrice:    s.LimitPrice,
			End:           t,
			Stop:          s.Stop,
			StopDirection: s.Side.toStopDirection(),
		},
	})
}

func (s *StopLimitOrderGTD) UnmarshalJSON(buf []byte) error {
	var w stopLimitGTDWrapper
	if err := json.Unmarshal(buf, &w); err != nil {
		return err
	}

	*s = StopLimitOrderGTD{
		BaseSize:   w.Data.BaseSize,
		LimitPrice: w.Data.LimitPrice,
		Stop:       w.Data.Stop,
		EndTime:    w.Data.end(),
		Side:       w.Data.stopSide(),
	}

	return nil
}
