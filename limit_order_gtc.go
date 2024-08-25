package coinbase

import "encoding/json"

type limitGtcWrapper struct {
	Data orderData `json:"limit_limit_gtc"`
}

type LimitOrderGTC struct {
	// The amount of the first Asset in the Trading Pair. For example, on the
	// BTC-USD Order Book, BTC is the Base Asset.
	BaseSize string

	// The specified price, or better, that the Order should be executed at. A
	// Buy Order will execute at or lower than the limit price. A Sell Order
	// will execute at or higher than the limit price.
	LimitPrice string

	// Enable or disable Post-only Mode. When enabled, only Maker Orders will be
	// posted to the Order Book. Orders that will be posted as a Taker Order
	// will be rejected.
	PostOnly bool
}

func (l *LimitOrderGTC) OrderType() OrderType { return LimitGTC }

func (l *LimitOrderGTC) MarshalJSON() ([]byte, error) {
	return json.Marshal(limitGtcWrapper{
		Data: orderData{
			BaseSize:   l.BaseSize,
			LimitPrice: l.LimitPrice,
			PostOnly:   l.PostOnly,
		},
	})
}

func (l *LimitOrderGTC) UnmarshalJSON(b []byte) error {
	var x limitGtcWrapper
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}

	*l = LimitOrderGTC{
		BaseSize:   x.Data.BaseSize,
		LimitPrice: x.Data.LimitPrice,
		PostOnly:   x.Data.PostOnly,
	}

	return nil
}
