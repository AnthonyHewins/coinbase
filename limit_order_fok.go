package coinbase

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type limitFOKWrapper struct {
	Data *orderData `json:"limit_limit_fok"`
}

type LimitOrderFOK struct {
	// The amount of the first Asset in the Trading Pair. For example, on the
	// BTC-USD Order Book, BTC is the Base Asset.
	BaseSize decimal.Decimal `json:"base_size"`

	// The specified price, or better, that the Order should be executed at. A
	// Buy Order will execute at or lower than the limit price. A Sell Order
	// will execute at or higher than the limit price.
	LimitPrice decimal.Decimal `json:"limit_price"`
}

func (l *LimitOrderFOK) OrderType() OrderType { return LimitFOK }

func (l *LimitOrderFOK) MarshalJSON() ([]byte, error) {
	return json.Marshal(limitFOKWrapper{
		Data: &orderData{
			BaseSize:   l.BaseSize,
			LimitPrice: l.LimitPrice,
		},
	})
}

func (l *LimitOrderFOK) UnmarshalJSON(b []byte) error {
	var x limitFOKWrapper
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}

	*l = LimitOrderFOK{
		BaseSize:   x.Data.BaseSize,
		LimitPrice: x.Data.LimitPrice,
	}

	return nil
}
