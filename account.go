package coinbase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

//go:generate enumer -type AcctType -transform snake_upper -json
type AcctType byte

const (
	AccountTypeUnspecified AcctType = iota
	AccountTypeCrypto
	AccountTypeFiat
	AccountTypeVault
	AccountTypePerpFutures
)

type Balance struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type Account struct {
	ID                uuid.UUID `json:"uuid"`
	Name              string    `json:"name"`
	Currency          string    `json:"currency"`
	AvailableBalance  Balance   `json:"available_balance"`
	Default           bool      `json:"default"`
	Active            bool      `json:"active"`
	Created           time.Time `json:"created_at"`
	Updated           time.Time `json:"updated_at"`
	Deleted           time.Time `json:"deleted_at"`
	Type              AcctType  `json:"type"`
	Ready             bool      `json:"ready"`
	Hold              Balance   `json:"hold"`
	RetailPortfolioID string    `json:"retail_portfolio_id"`
}

func (c *Client) ListAccounts(ctx context.Context) ([]Account, error) {
	type wrapper struct {
		Accts   []Account `json:"accounts"`
		HasNext bool      `json:"has_next"`
		Cursor  string    `json:"cursor"`
		Size    int32     `json:"size"`
	}

	var accounts wrapper
	err := c.get(ctx, "/accounts", nil, &accounts)
	return accounts.Accts, err
}

func (c *Client) GetAccount(ctx context.Context, id uuid.UUID) (*Account, error) {
	type wrapper struct {
		Acct Account `json:"account"`
	}

	w := wrapper{}
	url := fmt.Sprintf("/accounts/%s", id)
	_, err := c.request(ctx, "GET", url, nil, &w)
	return &w.Acct, err
}
