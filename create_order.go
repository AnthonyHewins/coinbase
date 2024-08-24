package coinbase

import "context"

func (c *Client) CreateOrder(ctx context.Context, newOrder *Order) (created bool, err error) {
	type createOrderResp struct {
		Success bool `json:"success"`
	}

	var savedOrder createOrderResp
	_, err = c.Request(ctx, "POST", "/orders", newOrder, &savedOrder)
	return savedOrder.Success, err
}
