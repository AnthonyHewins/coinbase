package coinbase

import "context"

type CreateOrderArgs struct {
	ID        string      `json:"client_order_id"`
	ProductID string      `json:"product_id"`
	Side      Side        `json:"side"`
	Config    OrderConfig `json:"order_configuration"`

	// Optional: empty string will omit leverage
	Leverage string `json:"leverage,omitempty"`

	// Optional: default value will not send
	MarginType MarginType `json:"margin_type,omitempty"`

	// Optional: empty string will omit
	RetailPortfolioID string `json:"retail_portfolio_id,omitempty"`

	// Optional: empty string will omit
	PreviewID string `json:"preview_id,omitempty"`
}

func (c *Client) CreateOrder(ctx context.Context, newOrder *CreateOrderArgs) (created bool, err error) {
	type createOrderResp struct {
		Success bool `json:"success"`
	}

	var savedOrder createOrderResp
	_, err = c.request(ctx, "POST", "/orders", newOrder, &savedOrder)
	return savedOrder.Success, err
}
