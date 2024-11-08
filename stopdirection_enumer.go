// Code generated by "enumer -type StopDirection -transform snake-upper -json"; DO NOT EDIT.

package coinbase

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _StopDirectionName = "STOP_DIRECTION_UNSPECIFIEDSTOP_DIRECTION_STOP_UPSTOP_DIRECTION_STOP_DOWN"

var _StopDirectionIndex = [...]uint8{0, 26, 48, 72}

const _StopDirectionLowerName = "stop_direction_unspecifiedstop_direction_stop_upstop_direction_stop_down"

func (i StopDirection) String() string {
	if i >= StopDirection(len(_StopDirectionIndex)-1) {
		return fmt.Sprintf("StopDirection(%d)", i)
	}
	return _StopDirectionName[_StopDirectionIndex[i]:_StopDirectionIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _StopDirectionNoOp() {
	var x [1]struct{}
	_ = x[StopDirectionUnspecified-(0)]
	_ = x[StopDirectionStopUp-(1)]
	_ = x[StopDirectionStopDown-(2)]
}

var _StopDirectionValues = []StopDirection{StopDirectionUnspecified, StopDirectionStopUp, StopDirectionStopDown}

var _StopDirectionNameToValueMap = map[string]StopDirection{
	_StopDirectionName[0:26]:       StopDirectionUnspecified,
	_StopDirectionLowerName[0:26]:  StopDirectionUnspecified,
	_StopDirectionName[26:48]:      StopDirectionStopUp,
	_StopDirectionLowerName[26:48]: StopDirectionStopUp,
	_StopDirectionName[48:72]:      StopDirectionStopDown,
	_StopDirectionLowerName[48:72]: StopDirectionStopDown,
}

var _StopDirectionNames = []string{
	_StopDirectionName[0:26],
	_StopDirectionName[26:48],
	_StopDirectionName[48:72],
}

// StopDirectionString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func StopDirectionString(s string) (StopDirection, error) {
	if val, ok := _StopDirectionNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _StopDirectionNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to StopDirection values", s)
}

// StopDirectionValues returns all values of the enum
func StopDirectionValues() []StopDirection {
	return _StopDirectionValues
}

// StopDirectionStrings returns a slice of all String values of the enum
func StopDirectionStrings() []string {
	strs := make([]string, len(_StopDirectionNames))
	copy(strs, _StopDirectionNames)
	return strs
}

// IsAStopDirection returns "true" if the value is listed in the enum definition. "false" otherwise
func (i StopDirection) IsAStopDirection() bool {
	for _, v := range _StopDirectionValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for StopDirection
func (i StopDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for StopDirection
func (i *StopDirection) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("StopDirection should be a string, got %s", data)
	}

	var err error
	*i, err = StopDirectionString(s)
	return err
}
