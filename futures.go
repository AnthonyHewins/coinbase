package coinbase

import (
	"time"

	"github.com/shopspring/decimal"
)

type FuturePosition struct {
	ProductId       string          `json:"product_id"`
	ContractSize    string          `json:"contract_size"`
	Side            string          `json:"side"`
	Amount          decimal.Decimal `json:"amount"`
	AvgEntryPrice   decimal.Decimal `json:"avg_entry_price"`
	CurrentPrice    decimal.Decimal `json:"current_price"`
	UnrealizedPnl   string          `json:"unrealized_pnl"`
	Expiry          time.Time       `json:"expiry"`
	UnderlyingAsset string          `json:"underlying_asset"`
	AssetImgUrl     string          `json:"asset_img_url"`
	ProductName     string          `json:"product_name"`
	Venue           string          `json:"venue"`
	NotionalValue   string          `json:"notional_value"`
}
