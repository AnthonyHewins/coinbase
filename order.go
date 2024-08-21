package coinbase

import (
	"bytes"
	"encoding/json"
	"fmt"
)

//go:generate enumer -type Side -json -trimprefix Side -transform upper
type Side byte

const (
	SideUnspecified Side = iota
	SideBuy
	SideSell
)

type OrderType byte

const (
	OrderTypeUnspecified OrderType = iota
	Market
	LimitIOC // Immediate or cancel (has to fill some part of the order immediately, if it fails, cancel the rest)
	LimitFOK // Fill or Kill
	LimitGTC // Good til Canceled
	LimitGTD // Good til date
	StopLimitGTC
	StopLimitGTD
	TriggerBracketGTC
	TriggerBracketGTD
)

type OrderConfig interface {
	orderType() OrderType
}

//go:generate enumer -type MarginType -json -transform upper -trimprefix MarginType
type MarginType byte

const (
	MarginTypeUnspecified MarginType = iota
	MarginTypeCross
	MarginTypeIsolated
)

type Order struct {
	ID        string      `json:"client_order_id"`
	ProductID string      `json:"product_id"`
	Side      Side        `json:"side"`
	Config    OrderConfig `json:"order_configuration"`

	// Optional: empty string will omit leverage
	Leverage string `json:"leverage"`

	// Optional: default value will not send
	MarginType MarginType `json:"margin_type"`

	// Optional: empty string will omit
	RetailPortfolioID string `json:"retail_portfolio_id"`

	// Optional: empty string will omit
	PreviewID string `json:"preview_id"`
	// Type      string `json:"type"`
	// Size      string `json:"size,omitempty"`
	// ProductID string `json:"product_id"`
	// ClientOID string `json:"client_oid,omitempty"`
	// Stp       string `json:"stp,omitempty"`
	// Stop      string `json:"stop,omitempty"`
	// StopPrice string `json:"stop_price,omitempty"`
	// // Limit Order
	// Price       string `json:"price,omitempty"`
	// TimeInForce string `json:"time_in_force,omitempty"`
	// PostOnly    bool   `json:"post_only,omitempty"`
	// CancelAfter string `json:"cancel_after,omitempty"`
	// // Market Order
	// Funds          string `json:"funds,omitempty"`
	// SpecifiedFunds string `json:"specified_funds,omitempty"`
	// // Response Fields
	// ID            string `json:"id"`
	// Status        string `json:"status,omitempty"`
	// Settled       bool   `json:"settled,omitempty"`
	// DoneReason    string `json:"done_reason,omitempty"`
	// DoneAt        Time   `json:"done_at,string,omitempty"`
	// CreatedAt     Time   `json:"created_at,string,omitempty"`
	// FillFees      string `json:"fill_fees,omitempty"`
	// FilledSize    string `json:"filled_size,omitempty"`
	// ExecutedValue string `json:"executed_value,omitempty"`
}

func (o *Order) MarshalJSON() ([]byte, error) {
	b := bytes.NewBuffer([]byte(`{"client_order_id":`))
	e := json.NewEncoder(b)
	e.SetIndent("", "")

	if err := e.Encode(o.ID); err != nil {
		return nil, err
	}
	b.WriteString(`,"product_id":`)

	if err := e.Encode(o.ProductID); err != nil {
		return nil, err
	}

	if o.Config != nil {
		b.WriteString(`,"order_configuration":`)
		if err := e.Encode(o.Config); err != nil {
			return nil, err
		}
	}

	type strFields struct {
		name, value string
	}

	for _, v := range []strFields{
		{"leverage", o.Leverage},
		{"retail_portfolio_id", o.RetailPortfolioID},
		{"preview_id", o.PreviewID},
	} {
		if v.value != "" {
			b.WriteString(fmt.Sprintf(`,"%s":`, v.name))
			if err := e.Encode(v.value); err != nil {
				return nil, err
			}
		}
	}

	if o.MarginType != MarginTypeUnspecified {
		b.WriteString(`,"margin_type":`)
		if err := e.Encode(o.MarginType); err != nil {
			return nil, err
		}
	}

	b.WriteRune('}')
	return b.Bytes(), nil
}

type CancelAllOrdersParams struct {
	ProductID string
}

type ListOrdersParams struct {
	Status     string
	ProductID  string
	Pagination PaginationParams
}

func (c *Client) CreateOrder(newOrder *Order) (Order, error) {
	var savedOrder Order

	url := fmt.Sprintf("/orders")
	_, err := c.Request("POST", url, newOrder, &savedOrder)
	return savedOrder, err
}

func (c *Client) CancelOrder(id string) error {
	url := fmt.Sprintf("/orders/%s", id)
	_, err := c.Request("DELETE", url, nil, nil)
	return err
}

func (c *Client) CancelAllOrders(p ...CancelAllOrdersParams) ([]string, error) {
	var orderIDs []string
	url := "/orders"

	if len(p) > 0 && p[0].ProductID != "" {
		url = fmt.Sprintf("%s?product_id=%s", url, p[0].ProductID)
	}

	_, err := c.Request("DELETE", url, nil, &orderIDs)
	return orderIDs, err
}

func (c *Client) GetOrder(id string) (Order, error) {
	var savedOrder Order

	url := fmt.Sprintf("/orders/%s", id)
	_, err := c.Request("GET", url, nil, &savedOrder)
	return savedOrder, err
}

func (c *Client) ListOrders(p ...ListOrdersParams) *Cursor {
	paginationParams := PaginationParams{}
	if len(p) > 0 {
		paginationParams = p[0].Pagination
		if p[0].Status != "" {
			paginationParams.AddExtraParam("status", p[0].Status)
		}
		if p[0].ProductID != "" {
			paginationParams.AddExtraParam("product_id", p[0].ProductID)
		}
	}

	return NewCursor(c, "GET", fmt.Sprintf("/orders"),
		&paginationParams)
}
