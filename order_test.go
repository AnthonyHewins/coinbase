package coinbase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCancelOrders(mainTest *testing.T) {
	testCases := []struct {
		name string

		mockResp cancelWrapper
		mockErr  *Error

		expectedErr string
	}{
		{
			name: "base case",
		},
		{
			name: "happy path",
			mockResp: cancelWrapper{
				Res: []cancelResponse{
					{
						ID:            "",
						Success:       true,
						FailureReason: "",
					},
				},
			},
		},
		{
			name: "failed cancel path",
			mockResp: cancelWrapper{
				Res: []cancelResponse{
					{ID: "x", FailureReason: "fail"},
					{ID: "y", FailureReason: "fail"},
				},
			},
			expectedErr: "fail: (Order ID: x)\nfail: (Order ID: y)",
		},
		{
			name: "error path",
			mockErr: &Error{
				Err:     "failed",
				Message: "s",
			},
			expectedErr: "failed: s",
		},
	}

	for _, tc := range testCases {
		mainTest.Run(tc.name, func(tt *testing.T) {
			t := assert.New(tt)

			var status int
			var x any
			if tc.mockErr != nil {
				x = tc.mockErr
				status = 500
			} else {
				x = tc.mockResp
				status = 200
			}

			s := newTestserver(status, x)
			defer s.server.Close()

			actualErr := s.c.CancelOrders(context.Background(), "")

			if tc.expectedErr != "" {
				t.EqualError(actualErr, tc.expectedErr)
				return
			}

			t.NoError(actualErr, "should not error when a value is expected")
		})
	}
}
