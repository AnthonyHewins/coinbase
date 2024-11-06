package coinbase

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type sorLimitIOCWrapper struct {
	Data *orderData `json:"sor_limit_ioc"`
}

type LimitOrderIOC struct {
	// The amount of the first Asset in the Trading Pair. For example, on the
	// BTC-USD Order Book, BTC is the Base Asset.
	BaseSize decimal.Decimal

	// The specified price, or better, that the Order should be executed at. A
	// Buy Order will execute at or lower than the limit price. A Sell Order
	// will execute at or higher than the limit price.
	LimitPrice decimal.Decimal
}

func (l *LimitOrderIOC) OrderType() OrderType { return LimitIOC }

func (l *LimitOrderIOC) MarshalJSON() ([]byte, error) {
	return json.Marshal(sorLimitIOCWrapper{
		Data: &orderData{
			BaseSize:   l.BaseSize,
			LimitPrice: l.LimitPrice,
		},
	})
}

func (l *LimitOrderIOC) UnmarshalJSON(b []byte) error {
	var x sorLimitIOCWrapper
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}

	*l = LimitOrderIOC{
		BaseSize:   x.Data.BaseSize,
		LimitPrice: x.Data.LimitPrice,
	}

	return nil
}
