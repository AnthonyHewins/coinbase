package coinbase

import (
	"context"
	"errors"
)

type cancelResponse struct {
	ID            string `json:"order_id"`
	Success       bool   `json:"success"`
	FailureReason string `json:"failure_reason"`
}

type cancelWrapper struct {
	Res []cancelResponse `json:"results"`
}

func (w cancelWrapper) errors() error {
	e := make([]error, len(w.Res))
	for i, v := range w.Res {
		e[i] = Error{
			Err:     v.FailureReason,
			Details: "Order ID: " + v.ID,
		}
	}

	return errors.Join(e...)
}

func (c *Client) CancelOrders(ctx context.Context, ids ...string) error {
	if len(ids) == 0 {
		return nil
	}

	var r cancelWrapper
	_, err := c.request(ctx, "DELETE", "/orders/batch_cancel", map[string][]string{
		"order_ids": ids,
	}, &r)

	if err != nil {
		return err
	}

	failedPtr := 0
	for _, v := range r.Res {
		if v.Success {
			continue
		}

		r.Res[failedPtr] = v
		failedPtr++
	}

	if failedPtr == 0 {
		return nil
	}

	return r.errors()
}
