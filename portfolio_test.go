package coinbase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/shopspring/decimal"
)

func TestSpotJSON(mainTest *testing.T) {
	testCases := []struct {
		arg any
	}{
		{
			arg: SpotPosition{
				Asset:                "asset",
				AccountUUID:          "uuid",
				TotalBalanceFiat:     1,
				TotalBalanceCrypto:   2,
				AvailableToTradeFiat: 3,
				Allocation:           4,
				OneDayChange:         5,
				Cost_basis: Balance{
					Value:    decimal.NewFromInt(1),
					Currency: "2",
				},
				AssetImgUrl: "asdin",
				IsCash:      true,
			},
		},
	}

	for _, tc := range testCases {
		mainTest.Run(fmt.Sprintf("%T", tc.arg), func(t *testing.T) {
			b := bytes.NewBuffer([]byte{})
			m := json.NewEncoder(b)
			m.SetIndent("\t", "")
			err := m.Encode(tc.arg)
			if err != nil {
				t.Errorf("should not fail marshal")
				return
			}

			snaps.MatchSnapshot(t, b.String())
		})
	}
}
