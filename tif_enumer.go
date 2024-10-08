// Code generated by "enumer -type TIF -json -transform snake_upper -trimprefix TIF"; DO NOT EDIT.

package coinbase

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _TIFName = "UNKNOWNGOOD_UNTIL_DATE_TIMEGOOD_UNTIL_CANCELLEDIMMEDIATE_OR_CANCELFILL_OR_KILL"

var _TIFIndex = [...]uint8{0, 7, 27, 47, 66, 78}

const _TIFLowerName = "unknowngood_until_date_timegood_until_cancelledimmediate_or_cancelfill_or_kill"

func (i TIF) String() string {
	if i >= TIF(len(_TIFIndex)-1) {
		return fmt.Sprintf("TIF(%d)", i)
	}
	return _TIFName[_TIFIndex[i]:_TIFIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _TIFNoOp() {
	var x [1]struct{}
	_ = x[TIFUnknown-(0)]
	_ = x[TIFGoodUntilDateTime-(1)]
	_ = x[TIFGoodUntilCancelled-(2)]
	_ = x[TIFImmediateOrCancel-(3)]
	_ = x[TIFFillOrKill-(4)]
}

var _TIFValues = []TIF{TIFUnknown, TIFGoodUntilDateTime, TIFGoodUntilCancelled, TIFImmediateOrCancel, TIFFillOrKill}

var _TIFNameToValueMap = map[string]TIF{
	_TIFName[0:7]:        TIFUnknown,
	_TIFLowerName[0:7]:   TIFUnknown,
	_TIFName[7:27]:       TIFGoodUntilDateTime,
	_TIFLowerName[7:27]:  TIFGoodUntilDateTime,
	_TIFName[27:47]:      TIFGoodUntilCancelled,
	_TIFLowerName[27:47]: TIFGoodUntilCancelled,
	_TIFName[47:66]:      TIFImmediateOrCancel,
	_TIFLowerName[47:66]: TIFImmediateOrCancel,
	_TIFName[66:78]:      TIFFillOrKill,
	_TIFLowerName[66:78]: TIFFillOrKill,
}

var _TIFNames = []string{
	_TIFName[0:7],
	_TIFName[7:27],
	_TIFName[27:47],
	_TIFName[47:66],
	_TIFName[66:78],
}

// TIFString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TIFString(s string) (TIF, error) {
	if val, ok := _TIFNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _TIFNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to TIF values", s)
}

// TIFValues returns all values of the enum
func TIFValues() []TIF {
	return _TIFValues
}

// TIFStrings returns a slice of all String values of the enum
func TIFStrings() []string {
	strs := make([]string, len(_TIFNames))
	copy(strs, _TIFNames)
	return strs
}

// IsATIF returns "true" if the value is listed in the enum definition. "false" otherwise
func (i TIF) IsATIF() bool {
	for _, v := range _TIFValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for TIF
func (i TIF) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for TIF
func (i *TIF) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("TIF should be a string, got %s", data)
	}

	var err error
	*i, err = TIFString(s)
	return err
}
