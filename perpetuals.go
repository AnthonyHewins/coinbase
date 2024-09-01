package coinbase

//go:generate enumer -type FuturesPosition -json -transform snake_upper
type FuturesPosition byte

const (
	FuturesPositionSideUnspecified FuturesPosition = iota
	FuturesPositionSideLong
	FuturesPositionSideShort
)

type Currency struct {
	UserNativeCurrency Balance `json:"userNativeCurrency"`
	RawCurrency        Balance `json:"rawCurrency"`
}

type PerpetualPosition struct {
	ProductId          string     `json:"product_id"`
	ProductUuid        string     `json:"product_uuid"`
	Symbol             string     `json:"symbol"`
	AssetImageUrl      string     `json:"asset_image_url"`
	NetSize            string     `json:"net_size"`
	BuyOrderSize       string     `json:"buy_order_size"`
	SellOrderSize      string     `json:"sell_order_size"`
	IMContribution     string     `json:"im_contribution"`
	VWAP               Currency   `json:"vwap"`
	UnrealizedPNL      Currency   `json:"unrealized_pnl"`
	MarkPrice          Currency   `json:"mark_price"`
	LiquidationPrice   Currency   `json:"liquidation_price"`
	Leverage           string     `json:"leverage"`
	IMNotational       Currency   `json:"im_notional"`
	MMNotational       Currency   `json:"mm_notional"`
	PositionNotational Currency   `json:"position_notional"`
	MarginType         MarginType `json:"margin_type"`
	LiquidationBuffer  string     `json:"liquidation_buffer"`
	LiquidationPct     string     `json:"liquidation_percentage"`
}
