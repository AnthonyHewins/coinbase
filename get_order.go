package coinbase

import (
	"context"
	"fmt"
)

type ListOrdersParams struct {
	Status     string
	ProductID  string
	Pagination PaginationParams
}

func (c *Client) GetOrder(ctx context.Context, id string) (Order, error) {
	var savedOrder Order

	url := fmt.Sprintf("/orders/%s", id)
	_, err := c.Request(ctx, "GET", url, nil, &savedOrder)
	return savedOrder, err
}
