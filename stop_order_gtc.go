package coinbase

import "encoding/json"

type stopLimitGTCWrapper struct {
	Data orderData `json:"stop_limit_stop_limit_gtc"`
}

type StopLimitOrderGTC struct {
	// The amount of the first Asset in the Trading Pair. For example, on the
	// BTC-USD Order Book, BTC is the Base Asset.
	BaseSize string

	// The specified price, or better, that the Order should be executed at. A
	// Buy Order will execute at or lower than the limit price. A Sell Order
	// will execute at or higher than the limit price.
	LimitPrice string
	Stop       string

	// on which side should the limit trigger.
	// if you specify Buy then the stop is triggered once the stop EXCEEDS.
	// if you specify Sell then the stop is triggered once the stop FALLS BELOW
	// If you don't specify one, the payload to coinbase will not have one,
	// and will result in an error
	Side Side
}

func (s *StopLimitOrderGTC) OrderType() OrderType { return StopLimitGTC }

func (s *StopLimitOrderGTC) MarshalJSON() ([]byte, error) {
	return json.Marshal(stopLimitGTCWrapper{
		Data: orderData{
			BaseSize:   s.BaseSize,
			LimitPrice: s.LimitPrice,
			Stop:       s.Stop,
			Side:       s.Side.toStopDirectionStr(),
		},
	})
}

func (s *StopLimitOrderGTC) UnmarshalJSON(buf []byte) error {
	var w stopLimitGTCWrapper
	if err := json.Unmarshal(buf, &w); err != nil {
		return err
	}

	side, err := w.Data.stopSide()
	if err != nil {
		return err
	}

	*s = StopLimitOrderGTC{
		BaseSize:   w.Data.BaseSize,
		LimitPrice: w.Data.LimitPrice,
		Stop:       w.Data.Stop,
		Side:       side,
	}

	return nil
}
