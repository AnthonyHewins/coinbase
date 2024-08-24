package coinbase

import (
	"fmt"
	"strings"
)

type Error struct {
	Err     string `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"error_details"`
}

func (e Error) Error() string {
	var sb strings.Builder

	if e.Code != 0 {
		sb.WriteString(fmt.Sprintf("%d ", e.Code))
	}

	if e.Err != "" {
		sb.WriteString(fmt.Sprintf("%s: ", e.Err))
	}

	sb.WriteString(e.Message)

	if e.Details != "" {
		sb.WriteString(fmt.Sprintf("(%s)", e.Details))
	}

	return sb.String()
}
