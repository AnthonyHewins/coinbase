package coinbase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type OrderBookEntry struct {
	Price string `json:"price"`
	Size  string `json:"size"`
}

type BidAsk struct {
	ProductID string           `json:"product_id"`
	Time      time.Time        `json:"time"`
	Bids      []OrderBookEntry `json:"bids"`
	Asks      []OrderBookEntry `json:"asks"`
}

func (c *Client) BidAsk(ctx context.Context, pairs ...string) ([]BidAsk, error) {
	type wrapper struct {
		Books []BidAsk `json:"pricebooks"`
	}

	params := url.Values{}
	for _, v := range pairs {
		params.Add("product_ids", v)
	}

	var w wrapper
	_, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/best_bid_ask?%s", params.Encode()), nil, &w)
	if err != nil {
		return nil, err
	}

	return w.Books, nil
}