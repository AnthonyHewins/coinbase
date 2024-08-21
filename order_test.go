package coinbase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderMarshal(mainTest *testing.T) {
	testCases := []struct {
		name        string
		arg         Order
		expected    string
		expectedErr error
	}{
		{
			name:        "base case",
			arg:         Order{},
			expected:    "",
			expectedErr: nil,
		},
	}

	t := assert.New(mainTest)
	for _, tc := range testCases {
		actual, actualErr := tc.arg.MarshalJSON()
		t.Equal(tc.expected, string(actual), tc.name)
		t.Equal(tc.expectedErr, actualErr, tc.name)
	}
}
