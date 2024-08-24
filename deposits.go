package coinbase

import (
	"context"
	"fmt"
)

type Deposit struct {
	Currency        string `json:"currency"`
	Amount          string `json:"amount"`
	PaymentMethodID string `json:"payment_method_id"` // PaymentMethodID can be determined by calling GetPaymentMethods()
	// Response fields
	ID       string `json:"id,omitempty"`
	PayoutAt Time   `json:"payout_at,string,omitempty"`
}

func (c *Client) CreateDeposit(ctx context.Context, newDeposit *Deposit) (Deposit, error) {
	var savedDeposit Deposit

	url := fmt.Sprintf("/deposits/payment-method")
	_, err := c.request(ctx, "POST", url, newDeposit, &savedDeposit)
	return savedDeposit, err
}

type PaymentMethod struct {
	Currency string `json:"currency"`
	Type     string `json:"type"`
	ID       string `json:"id"`
}

func (c *Client) GetPaymentMethods(ctx context.Context) ([]PaymentMethod, error) {
	var paymentMethods []PaymentMethod

	url := fmt.Sprintf("/payment-methods")
	_, err := c.request(ctx, "GET", url, nil, &paymentMethods)

	return paymentMethods, err
}
