package coinbase

import "encoding/json"

type limitFOKWrapper struct {
	Data orderData `json:"limit_limit_fok"`
}

type LimitOrderFOK struct {
	BaseSize   string `json:"base_size"`
	LimitPrice string `json:"limit_price"`
}

func (l *LimitOrderFOK) OrderType() OrderType { return LimitFOK }

func (l *LimitOrderFOK) MarshalJSON() ([]byte, error) {
	return json.Marshal(limitFOKWrapper{
		Data: orderData{
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
