package coinbase

import (
	"context"
	"fmt"
)

type Fees struct {
	MakerFeeRate string `json:"maker_fee_rate"`
	TakerFeeRate string `json:"taker_fee_rate"`
	USDVolume    string `json:"usd_volume"`
}

func (c *Client) GetFees(ctx context.Context) (Fees, error) {
	var fees Fees

	url := fmt.Sprintf("/fees")
	_, err := c.request(ctx, "GET", url, nil, &fees)
	return fees, err
}
