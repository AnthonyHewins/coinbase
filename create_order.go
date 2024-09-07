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

type createOrderErrResp struct {
	Err                   string `json:"error"` // undocumented field
	Msg                   string `json:"message"`
	Details               string `json:"error_details"`
	ReasonThatsDeprecated string `json:"preview_failure_reason"`   // supposed to be deprecated, still used
	ReasonThatShouldWork  string `json:"new_order_failure_reason"` // supposed to not be deprecated, not used as of now...ugh
}

func (c *createOrderErrResp) toErr() error {
	if c.Err == "" {
		return nil
	}

	for _, v := range []string{c.ReasonThatShouldWork, c.ReasonThatsDeprecated} {
		if v != "" {
			c.Msg = v
		}
	}

	return &Error{
		Code:    200,
		Err:     c.Err,
		Message: c.Msg,
		Details: c.Details,
	}
}

func (c *Client) CreateOrder(ctx context.Context, newOrder *CreateOrderArgs) (*CreateOrderResp, error) {
	type wrapper struct {
		Success     bool `json:"success"`
		SuccessResp struct {
			ID uuid.UUID `json:"order_id"`
		} `json:"success_response"`
		ErrResp createOrderErrResp `json:"error_response"`
	}

	var savedOrder wrapper
	_, err := c.request(ctx, "POST", "/orders", newOrder, &savedOrder)
	if err != nil {
		return nil, err
	}

	// create order is a horrible API, there's a large amount of unnecessary
	// work needed to see if something's wrong because things that are marked deprecated
	// are still used, and new things aren't, loop over them. On top of that, it's returned on a 200
	if err = savedOrder.ErrResp.toErr(); err != nil {
		return nil, err
	}

	return &CreateOrderResp{
		ID:         savedOrder.SuccessResp.ID,
		WasCreated: savedOrder.Success,
	}, nil
}
