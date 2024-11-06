package coinbase

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Fill struct {
	TradeID   int             `json:"trade_id"`
	ProductID string          `json:"product_id"`
	Price     decimal.Decimal `json:"price"`
	Size      string          `json:"size"`
	FillID    string          `json:"order_id"`
	CreatedAt Time            `json:"created_at"`
	Fee       string          `json:"fee"`
	Settled   bool            `json:"settled"`
	Side      string          `json:"side"`
	Liquidity string          `json:"liquidity"`
}

type ListFillsParams struct {
	OrderID    string
	ProductID  string
	Pagination PaginationParams
}

func (c *Client) ListFills(p ListFillsParams) *Cursor {
	paginationParams := p.Pagination
	if p.OrderID != "" {
		paginationParams.AddExtraParam("order_id", p.OrderID)
	}
	if p.ProductID != "" {
		paginationParams.AddExtraParam("product_id", p.ProductID)
	}

	return NewCursor(c, "GET", fmt.Sprintf("/fills"),
		&paginationParams)
}
