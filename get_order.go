package coinbase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type ListOrdersParams struct {
	Status     string
	ProductID  string
	Pagination PaginationParams
}

func (c *Client) GetOrder(ctx context.Context, id uuid.UUID) (*Order, error) {
	type wrapper struct {
		Order Order `json:"order"`
	}

	var w wrapper
	url := fmt.Sprintf("/orders/historical/%s", id)
	_, err := c.request(ctx, "GET", url, nil, &w)
	return &w.Order, err
}
