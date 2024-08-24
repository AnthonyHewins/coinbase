package coinbase

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type EditHistory struct {
	Price                  string
	Size                   string
	ReplaceAcceptTimestamp time.Time
}

type OrderResp struct {
	ID                    string // Coinbase's ID
	IdemKey               string // idempotency key you used to create it
	Product               string
	User                  string
	Config                OrderConfig
	Side                  Side
	Status                Status
	TIF                   TIF // Time in force
	Created               time.Time
	Completion            string // amount of the order that's been filled
	FilledSize            string
	AvgFillPrice          string
	NumberOfFills         string
	FilledValue           string
	PendingCancel         bool
	SizeInQuote           bool
	TotalFees             string
	SizeInclusiveOfFees   bool
	TotalValueAfterFees   string
	TriggerStatus         TriggerStatus
	RejectReason          RejectReason
	Settled               bool
	ProductType           ProductType
	RejectMsg             string
	CancelMsg             string
	OrderPlacementSrc     OrderPlacementSrc
	OutstandingHoldAmount string
	Liquidation           bool
	LastFillTime          time.Time
	Edits                 []EditHistory
	Leverage              string
	MarginType            MarginType
	RetailPortfolioID     string
}

func (o *OrderResp) UnmarshalJSON(b []byte) error {
	type wrapper struct {
		ID                    string            `json:"order_id"`
		IdemKey               string            `json:"client_order_id"`
		Product               string            `json:"product_id"`
		User                  string            `json:"user_id"`
		Config                json.RawMessage   `json:"order_configuration"`
		Side                  Side              `json:"side"`
		Status                Status            `json:"status"`
		TIF                   TIF               `json:"time_in_force"`
		Created               time.Time         `json:"created_time"`
		Completion            string            `json:"completion_percentage"`
		FilledSize            string            `json:"filled_size"`
		AvgFillPrice          string            `json:"average_filled_price"`
		NumberOfFills         string            `json:"number_of_fills"`
		FilledValue           string            `json:"filled_value"`
		PendingCancel         bool              `json:"pending_cancel"`
		SizeInQuote           bool              `json:"size_in_quote"`
		TotalFees             string            `json:"total_fees"`
		SizeInclusiveOfFees   bool              `json:"size_inclusive_of_fees"`
		TotalValueAfterFees   string            `json:"total_value_after_fees"`
		TriggerStatus         TriggerStatus     `json:"trigger_status"`
		OrderType             string            `json:"order_type"`
		RejectReason          RejectReason      `json:"reject_reason"`
		Settled               bool              `json:"settled"`
		ProductType           ProductType       `json:"product_type"`
		RejectMsg             string            `json:"reject_message"`
		CancelMsg             string            `json:"cancel_message"`
		OrderPlacementSrc     OrderPlacementSrc `json:"order_placement_src"`
		OutstandingHoldAmount string            `json:"outstanding_hold_amount"`
		Liquidation           bool              `json:"is_liquidation"`
		LastFillTime          time.Time         `json:"last_fill_time"`
		Edits                 []EditHistory     `json:"edit_history"`
		Leverage              string            `json:"leverage"`
		MarginType            MarginType        `json:"margin_type"`
		RetailPortfolioID     string            `json:"retail_portfolio_id"`
	}

	var x wrapper
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}

	conf, err := discoverConfig(x.OrderType, x.Config)
	if err != nil {
		return err
	}

	*o = OrderResp{
		ID:                    x.ID,
		IdemKey:               x.IdemKey,
		Product:               x.Product,
		User:                  x.User,
		Config:                conf,
		Side:                  x.Side,
		Status:                x.Status,
		TIF:                   x.TIF,
		Created:               x.Created,
		Completion:            x.Completion,
		FilledSize:            x.FilledSize,
		AvgFillPrice:          x.AvgFillPrice,
		NumberOfFills:         x.NumberOfFills,
		FilledValue:           x.FilledSize,
		PendingCancel:         x.PendingCancel,
		SizeInQuote:           x.SizeInQuote,
		TotalFees:             x.TotalFees,
		SizeInclusiveOfFees:   x.SizeInclusiveOfFees,
		TotalValueAfterFees:   x.TotalValueAfterFees,
		TriggerStatus:         x.TriggerStatus,
		RejectReason:          x.RejectReason,
		Settled:               x.Settled,
		ProductType:           x.ProductType,
		RejectMsg:             x.RejectMsg,
		CancelMsg:             x.CancelMsg,
		OrderPlacementSrc:     x.OrderPlacementSrc,
		OutstandingHoldAmount: x.OutstandingHoldAmount,
		Liquidation:           x.Liquidation,
		LastFillTime:          x.LastFillTime,
		Edits:                 x.Edits,
		Leverage:              x.Leverage,
		MarginType:            x.MarginType,
		RetailPortfolioID:     x.RetailPortfolioID,
	}

	return nil
}

type ListOrdersResp struct {
	Orders  []Order
	HasNext bool
}

func (c *Client) ListOrders(ctx context.Context) (*ListOrdersResp, error) {
	var r ListOrdersResp
	_, err := c.Request(ctx, http.MethodGet, "/rders/historical/batch", nil, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
