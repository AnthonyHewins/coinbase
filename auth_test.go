package coinbase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockNonce struct{}

func (m mockNonce) Nonce() (string, error) {
	return "", nil
}

func TestAuth(mainTest *testing.T) {
	testCases := []struct {
		name        string
		arg         Client
		expectedErr string
	}{
		{
			name:        "base case",
			expectedErr: "jwt: Could not decode private key (key length: 0)",
		},
		{
			name: "works with mock key",
			arg: Client{
				keyName:   "name",
				keySecret: mockPK,
			},
		},
	}

	t := assert.New(mainTest)
	for _, tc := range testCases {
		actual, actualErr := tc.arg.generateToken("method", "path")

		if tc.expectedErr != "" {
			if t.Empty(actual, "should not return a token if error is present") {
				t.EqualError(actualErr, tc.expectedErr)
			}
			continue
		}

		if t.NoError(actualErr) {
			t.NotEmpty(actual, tc.name)
		}
	}
}
