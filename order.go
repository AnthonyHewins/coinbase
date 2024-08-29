package coinbase

import "encoding/json"

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

func (s Side) toStopDirectionStr() string {
	switch s {
	case SideBuy:
		return "STOP_DIRECTION_STOP_UP"
	case SideSell:
		return "STOP_DIRECTION_STOP_DOWN"
	default:
		return ""
	}
}

//go:generate enumer -type MarginType -json -transform upper -trimprefix MarginType
type MarginType byte

const (
	MarginTypeUnspecified MarginType = iota
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
	TIFGoodUntilCanceled
	TIFImmediateOrCancel
	TIFFillOrKill
)

//go:generate enumer -type TriggerStatus -transform snake_upper -trimprefix TriggerStatus -json
type TriggerStatus byte

const (
	UnknownTriggerStatus TriggerStatus = iota
	TriggerStatusInvalidOrderType
	TriggerStatusInvalidStopPending
	TriggerStatusInvalidStopTriggered
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
