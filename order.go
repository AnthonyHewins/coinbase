package coinbase

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

//go:generate enumer -type OrderType
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
	json.Marshaler
	json.Unmarshaler
	OrderType() OrderType
}

//go:generate enumer -type Side -json -trimprefix Side -transform upper
type Side byte

const (
	SideUnspecified Side = iota
	SideBuy
	SideSell
)

func (s Side) toStopDirection() StopDirection {
	switch s {
	case SideBuy:
		return StopDirectionStopUp
	case SideSell:
		return StopDirectionStopDown
	default:
		return StopDirectionUnspecified
	}
}

//go:generate enumer -type MarginType -json -transform snake_upper -trimprefix MarginType
type MarginType byte

const (
	UnknownMarginType MarginType = iota
	MarginTypeCross
	MarginTypeIsolated
)

//go:generate enumer -type Status -json -transform upper -trimprefix Status
type Status byte

const (
	StatusUnknown Status = iota
	StatusPending
	StatusOpen
	StatusFilled
	StatusCancelled
	StatusExpired
	StatusFailed
	StatusQueued
	StatusCancelQueued
)

//go:generate enumer -type TIF -json -transform snake_upper -trimprefix TIF
type TIF byte

const (
	TIFUnknown TIF = iota
	TIFGoodUntilDateTime
	TIFGoodUntilCancelled
	TIFImmediateOrCancel
	TIFFillOrKill
)

//go:generate enumer -type TriggerStatus -transform snake_upper -trimprefix TriggerStatus -json
type TriggerStatus byte

const (
	UnknownTriggerStatus TriggerStatus = iota
	TriggerStatusInvalidOrderType
	TriggerStatusStopPending
	TriggerStatusStopTriggered
)

//go:generate enumer -type RejectReason -transform snake_upper -json
type RejectReason byte

const (
	// These values vary wildly for some reason. The API in coinbase for this is
	// not consistent
	RejectReasonUnspecified RejectReason = iota
	HoldFailure
	TooManyOpenOrders
	RejectReasonInsufficientFunds
	RateLimitExceeded
)

//go:generate enumer -type ProductType -transform snake_upper -json -trimprefix ProductType
type ProductType byte

const (
	UnknownProductType ProductType = iota
	ProductTypeSpot
	ProductTypeFuture
)

//go:generate enumer -type OrderPlacementSrc -transform snake_upper -json -trimprefix OrderPlacementSrc
type OrderPlacementSrc byte

const (
	UnknownPlacementSource OrderPlacementSrc = iota
	OrderPlacementSrcRetailSimple
	OrderPlacementSrcRetailAdvanced
)

type Order struct {
	ID                    uuid.UUID // Coinbase's ID
	IdemKey               string    // idempotency key you used to create it
	ProductID             string
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

func (o *Order) UnmarshalJSON(b []byte) error {
	type wrapper struct {
		ID                    uuid.UUID         `json:"order_id"`
		Product               string            `json:"product_id"`
		User                  string            `json:"user_id"`
		Config                json.RawMessage   `json:"order_configuration"`
		Side                  Side              `json:"side"`
		IdemKey               string            `json:"client_order_id"`
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
		OrderPlacementSrc     OrderPlacementSrc `json:"order_placement_source"`
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

	if len(x.Config) == 0 {
		return fmt.Errorf("order configuration is missing")
	}

	conf, err := discoverConfig(x.OrderType, x.Config)
	if err != nil {
		return err
	}

	*o = Order{
		ID:                    x.ID,
		IdemKey:               x.IdemKey,
		ProductID:             x.Product,
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
