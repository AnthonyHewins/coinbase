package coinbase

import (
	"context"

	"github.com/google/uuid"
)

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

type CreateOrderResp struct {
	ID         uuid.UUID
	WasCreated bool
}

func (c *Client) CreateOrder(ctx context.Context, newOrder *CreateOrderArgs) (*CreateOrderResp, error) {
	type wrapper struct {
		Success     bool `json:"success"`
		SuccessResp struct {
			ID uuid.UUID `json:"order_id"`
		} `json:"success_response"`
	}

	var savedOrder wrapper
	_, err := c.request(ctx, "POST", "/orders", newOrder, &savedOrder)
	if err != nil {
		return nil, err
	}

	return &CreateOrderResp{
		ID:         savedOrder.SuccessResp.ID,
		WasCreated: savedOrder.Success,
	}, nil
}
