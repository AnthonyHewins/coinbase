package coinbase

import (
	"context"
	"net/http"
	"time"
)

type EditHistory struct {
	Price                  string
	Size                   string
	ReplaceAcceptTimestamp time.Time
}

type ListOrdersResp struct {
	Orders  []Order `json:"orders"`
	HasNext bool    `json:"has_next"`
}

func (c *Client) ListOrders(ctx context.Context) (*ListOrdersResp, error) {
	var r ListOrdersResp
	_, err := c.request(ctx, http.MethodGet, "/orders/historical/batch", nil, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
