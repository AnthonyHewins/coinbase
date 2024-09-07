package tests

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/AnthonyHewins/coinbase"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestOrders(mainTest *testing.T) {
	if !riskyMode {
		mainTest.Skip("skipping riskier test")
		return
	}

	c := testClient()
	t := assert.New(mainTest)

	createResp, err := c.CreateOrder(context.Background(), &coinbase.CreateOrderArgs{
		ID:        "integration-test-order-" + uuid.NewString(),
		ProductID: "BTC-USD",
		Side:      coinbase.SideBuy,
		Config: &coinbase.LimitOrderGTD{
			BaseSize:   "0.00000001",
			LimitPrice: "0.01",
			EndTime:    time.Now().Add(time.Second * 15),
		},
	})

	if !t.NoError(err, "should not error creating order") {
		return
	}

	if !t.True(createResp.WasCreated, "order creation should have been executed") {
		return
	}

	defer func() {
		err = c.CancelOrders(context.Background(), createResp.ID.String())
		t.NoError(err, "canceling order should not fail")
	}()

	time.Sleep(2 * time.Second) // make sure the order gets posted

	orders, err := c.ListOrders(context.Background())
	if !t.NoError(err, "should not fail listing orders") {
		return
	}

	found := false
	for _, v := range orders.Orders {
		if found = v.ID == createResp.ID; found {
			break
		}
	}

	if !t.True(found, "order %s should be found after creation when listing orders", createResp.ID) {
		log.Printf("Got: %+v\n", orders.Orders)
		return
	}

	_, err = c.GetOrder(context.Background(), createResp.ID)
	if !t.NoError(err, "should not fail getting order %s", createResp.ID) {
		return
	}
}

func TestOrderFailure(tt *testing.T) {
	c := testClient()
	t := assert.New(tt)

	_, err := c.CreateOrder(context.Background(), &coinbase.CreateOrderArgs{
		ID:        "integration-test-order-" + uuid.NewString(),
		ProductID: "BTC-USD",
		Side:      coinbase.SideBuy,
		Config: &coinbase.LimitOrderGTD{
			BaseSize:   "0",
			LimitPrice: "0.01",
			EndTime:    time.Now().Add(time.Second),
		},
	})

	if t.EqualError(err, "200 UNSUPPORTED_ORDER_CONFIGURATION: PREVIEW_INVALID_ORDER_CONFIG", "should not error creating order") {
		return
	}
}
