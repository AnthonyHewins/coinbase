package coinbase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

//go:generate enumer -type PortfolioType -transform snake_upper -trimprefix PortfolioType -json
type PortfolioType byte

const (
	PortfolioTypeUndefined PortfolioType = iota
	PortfolioTypeDefault
	PortfolioTypeConsumer
	PortfolioTypeIntx
)

type PortfolioView struct {
	Name    string        `json:"name"`
	UUID    uuid.UUID     `json:"uuid"`
	Type    PortfolioType `json:"type"`
	Deleted bool          `json:"deleted"`
}

type SpotPosition struct {
	Asset                string  `json:"asset"`
	AccountUUID          string  `json:"account_uuid"`
	TotalBalanceFiat     float64 `json:"total_balance_fiat"`
	TotalBalanceCrypto   float64 `json:"total_balance_crypto"`
	AvailableToTradeFiat float64 `json:"available_to_trade_fiat"`
	Allocation           float64 `json:"allocation"`
	OneDayChange         float64 `json:"one_day_change"`
	Cost_basis           Balance `json:"cost_basis"`
	AssetImgUrl          string  `json:"asset_img_url"`
	IsCash               bool    `json:"is_cash"`
}

type TotalBalances struct {
	TotalBalance      Balance `json:"total_balance"`
	Futures           Balance `json:"total_futures_balance"`
	CashEq            Balance `json:"total_cash_equivalent_balance"`
	Crypto            Balance `json:"total_crypto_balance"`
	FuturesUnrealized Balance `json:"futures_unrealized_pnl"`
	PerpUnrealized    Balance `json:"perp_unrealized_pnl"`
}

type Portfolio struct {
	PortfolioView PortfolioView
	Balances      TotalBalances
	SpotPositions []SpotPosition
	Perpetuals    []PerpetualPosition
	Futures       []FuturePosition
}

func (c *Client) ListPortfolios(ctx context.Context) ([]PortfolioView, error) {
	type wrapper struct {
		Portfolios []PortfolioView `json:"portfolios"`
	}

	var w wrapper
	if err := c.get(ctx, "/portfolios", nil, &w); err != nil {
		return nil, err
	}

	return w.Portfolios, nil
}

func (c *Client) GetPortfolio(ctx context.Context, id uuid.UUID) (*Portfolio, error) {
	type wrapper struct {
		Breakdown struct {
			View      PortfolioView       `json:"portfolio"`
			Balances  TotalBalances       `json:"portfolio_balances"`
			Positions []SpotPosition      `json:"spot_positions"`
			Futures   []FuturePosition    `json:"futures_positions"`
			Perps     []PerpetualPosition `json:"perp_positions"`
		} `json:"breakdown"`
	}

	var w wrapper
	err := c.get(ctx, fmt.Sprintf("/portfolios/%s", id), nil, &w)
	if err != nil {
		return nil, err
	}

	return &Portfolio{
		PortfolioView: w.Breakdown.View,
		Balances:      w.Breakdown.Balances,
		SpotPositions: w.Breakdown.Positions,
		Perpetuals:    w.Breakdown.Perps,
		Futures:       w.Breakdown.Futures,
	}, nil
}
