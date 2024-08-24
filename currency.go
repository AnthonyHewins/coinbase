package coinbase

import (
	"context"
	"fmt"
)

type Currency struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	MinSize string `json:"min_size"`
}

func (c *Client) GetCurrencies(ctx context.Context) ([]Currency, error) {
	var currencies []Currency

	url := fmt.Sprintf("/currencies")
	_, err := c.request(ctx, "GET", url, nil, &currencies)
	return currencies, err
}
