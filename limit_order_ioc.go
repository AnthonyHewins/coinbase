package coinbase

import "encoding/json"

type sorLimitIOCWrapper struct {
	Data orderData `json:"sor_limit_ioc"`
}

type LimitOrderIOC struct {
	BaseSize   string
	LimitPrice string
}

func (l *LimitOrderIOC) OrderType() OrderType { return LimitIOC }

func (l *LimitOrderIOC) MarshalJSON() ([]byte, error) {
	return json.Marshal(sorLimitIOCWrapper{
		Data: orderData{
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
