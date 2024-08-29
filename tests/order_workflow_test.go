package tests

import (
	"context"
	"testing"
	"time"

	"github.com/AnthonyHewins/coinbase"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func equalOrder(t *assert.Assertions, c *coinbase.CreateOrderArgs, o *coinbase.Order) {
	t.Equal(
		&coinbase.CreateOrderArgs{
			ID:                c.ID,
			ProductID:         c.ProductID,
			Side:              c.Side,
			Config:            c.Config,
			Leverage:          "1", // leverage should default to 1
			MarginType:        c.MarginType,
			RetailPortfolioID: c.RetailPortfolioID,
			PreviewID:         c.PreviewID,
		},
		&coinbase.CreateOrderArgs{
			ID:                o.IdemKey,
			ProductID:         o.ProductID,
			Side:              o.Side,
			Leverage:          o.Leverage, // leverage should default to 1
			MarginType:        o.MarginType,
			RetailPortfolioID: o.RetailPortfolioID,
			Config:            c.Config,
		},
	)
}

func TestOrders(mainTest *testing.T) {
	if !riskyMode {
		mainTest.Skip("skipping riskier test")
		return
	}

	c := testClient()
	t := assert.New(mainTest)

	create := &coinbase.CreateOrderArgs{
		ID:        uuid.New().String(),
		ProductID: "BTC-USD",
		Side:      coinbase.SideBuy,
		Config: &coinbase.LimitOrderGTD{
			BaseSize:   "0.00000001",
			LimitPrice: "0.01",
			EndTime:    time.Now().Add(time.Second * 10),
		},
	}

	created, err := c.CreateOrder(context.Background(), create)

	if !t.NoError(err, "should not error creating order") {
		return
	}

	if !t.True(created, "order creation should have been executed") {
		return
	}

	order, err := c.GetOrder(context.Background(), create.ID)
	if !t.NoError(err, "should not fail getting order") {
		return
	}

	equalOrder(t, create, order)

	err = c.CancelOrders(context.Background(), create.ID)

	t.NoError(err, "canceling order should not fail")
}
