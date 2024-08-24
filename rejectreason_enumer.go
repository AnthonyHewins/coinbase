// Code generated by "enumer -type RejectReason -transform snake_upper -json"; DO NOT EDIT.

package coinbase

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _RejectReasonName = "REJECT_REASON_UNSPECIFIEDHOLD_FAILURETOO_MANY_OPEN_ORDERSREJECT_REASON_INSUFFICIENT_FUNDSRATE_LIMIT_EXCEEDED"

var _RejectReasonIndex = [...]uint8{0, 25, 37, 57, 89, 108}

const _RejectReasonLowerName = "reject_reason_unspecifiedhold_failuretoo_many_open_ordersreject_reason_insufficient_fundsrate_limit_exceeded"

func (i RejectReason) String() string {
	if i >= RejectReason(len(_RejectReasonIndex)-1) {
		return fmt.Sprintf("RejectReason(%d)", i)
	}
	return _RejectReasonName[_RejectReasonIndex[i]:_RejectReasonIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _RejectReasonNoOp() {
	var x [1]struct{}
	_ = x[RejectReasonUnspecified-(0)]
	_ = x[HoldFailure-(1)]
	_ = x[TooManyOpenOrders-(2)]
	_ = x[RejectReasonInsufficientFunds-(3)]
	_ = x[RateLimitExceeded-(4)]
}

var _RejectReasonValues = []RejectReason{RejectReasonUnspecified, HoldFailure, TooManyOpenOrders, RejectReasonInsufficientFunds, RateLimitExceeded}

var _RejectReasonNameToValueMap = map[string]RejectReason{
	_RejectReasonName[0:25]:        RejectReasonUnspecified,
	_RejectReasonLowerName[0:25]:   RejectReasonUnspecified,
	_RejectReasonName[25:37]:       HoldFailure,
	_RejectReasonLowerName[25:37]:  HoldFailure,
	_RejectReasonName[37:57]:       TooManyOpenOrders,
	_RejectReasonLowerName[37:57]:  TooManyOpenOrders,
	_RejectReasonName[57:89]:       RejectReasonInsufficientFunds,
	_RejectReasonLowerName[57:89]:  RejectReasonInsufficientFunds,
	_RejectReasonName[89:108]:      RateLimitExceeded,
	_RejectReasonLowerName[89:108]: RateLimitExceeded,
}

var _RejectReasonNames = []string{
	_RejectReasonName[0:25],
	_RejectReasonName[25:37],
	_RejectReasonName[37:57],
	_RejectReasonName[57:89],
	_RejectReasonName[89:108],
}

// RejectReasonString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func RejectReasonString(s string) (RejectReason, error) {
	if val, ok := _RejectReasonNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _RejectReasonNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to RejectReason values", s)
}

// RejectReasonValues returns all values of the enum
func RejectReasonValues() []RejectReason {
	return _RejectReasonValues
}

// RejectReasonStrings returns a slice of all String values of the enum
func RejectReasonStrings() []string {
	strs := make([]string, len(_RejectReasonNames))
	copy(strs, _RejectReasonNames)
	return strs
}

// IsARejectReason returns "true" if the value is listed in the enum definition. "false" otherwise
func (i RejectReason) IsARejectReason() bool {
	for _, v := range _RejectReasonValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for RejectReason
func (i RejectReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for RejectReason
func (i *RejectReason) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("RejectReason should be a string, got %s", data)
	}

	var err error
	*i, err = RejectReasonString(s)
	return err
}
