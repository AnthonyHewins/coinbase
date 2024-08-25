package coinbase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID `json:"id"`
	Balance   string    `json:"balance"`
	Hold      string    `json:"hold"`
	Available string    `json:"available"`
	Currency  string    `json:"currency"`
}

type LedgerEntry struct {
	ID        string        `json:"id"`
	CreatedAt Time          `json:"created_at"`
	Amount    string        `json:"amount"`
	Balance   string        `json:"balance"`
	Type      string        `json:"type"`
	Details   LedgerDetails `json:"details"`
}

type LedgerDetails struct {
	OrderID   string `json:"order_id"`
	TradeID   string `json:"trade_id"`
	ProductID string `json:"product_id"`
}

type GetAccountLedgerParams struct {
	Pagination PaginationParams
}

type Hold struct {
	AccountID string `json:"account_id"`
	CreatedAt Time   `json:"created_at,string"`
	UpdatedAt Time   `json:"updated_at,string"`
	Amount    string `json:"amount"`
	Type      string `json:"type"`
	Ref       string `json:"ref"`
}

type ListHoldsParams struct {
	Pagination PaginationParams
}

// Client Funcs
func (c *Client) GetAccounts(ctx context.Context) ([]Account, error) {
	var accounts []Account
	_, err := c.request(ctx, "GET", "/accounts", nil, &accounts)

	return accounts, err
}

func (c *Client) GetAccount(ctx context.Context, id uuid.UUID) (Account, error) {
	account := Account{}

	url := fmt.Sprintf("/accounts/%s", id)
	_, err := c.request(ctx, "GET", url, nil, &account)
	return account, err
}

func (c *Client) ListAccountLedger(id string,
	p ...GetAccountLedgerParams) *Cursor {
	paginationParams := PaginationParams{}
	if len(p) > 0 {
		paginationParams = p[0].Pagination
	}

	return NewCursor(c, "GET", fmt.Sprintf("/accounts/%s/ledger", id),
		&paginationParams)
}

func (c *Client) ListHolds(id string, p ...ListHoldsParams) *Cursor {
	paginationParams := PaginationParams{}
	if len(p) > 0 {
		paginationParams = p[0].Pagination
	}

	return NewCursor(c, "GET", fmt.Sprintf("/accounts/%s/holds", id),
		&paginationParams)
}
