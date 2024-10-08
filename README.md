Go Coinbase Pro [![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/AnthonyHewins/coinbase)
========

v3 [CoinBase Pro](https://pro.coinbase.com) API client fork, though it has basically zero
resemblance to the origin

- [Go Coinbase Pro ](#go-coinbase-pro-)
  - [Go get](#go-get)
  - [Decimal management](#decimal-management)
  - [Examples](#examples)
    - [Create order](#create-order)
      - [Cancel order(s)](#cancel-orders)
      - [List orders](#list-orders)

## Go get

```sh
go get github.com/AnthonyHewins/coinbase
```

## Decimal management

To manage precision correctly, this library sends all price values as strings for now.
Considering building a decimal library directly into the codebase, but not sure which one

## Examples

### Create order

All order types are available. To specify which one you want you need to pick
an implementation of the `OrderConfig` interface:

- `MarketOrder`
- `LimitOrderIOC`
- `LimitOrderFOK`
- `LimitOrderGTC`
- `LimitOrderGTD`
- `StopLimitOrderGTC`
- `StopLimitOrderGTD`
- `TriggerBracketOrderGTC`
- `TriggerBracketOrderGTD`

```go
idemKey := uuid.New() // idempotency ID you create
wasCreated, err := client.CreateOrder(ctx, &coinbase.Order{
  ID: idemKey.String(),
  ProductID: "BTC-USD",
  Side: coinbase.SideBuy,
  Leverage: "2",
  MarginType: MarginTypeCross,
  RetailPortfolioID: "123153432",
  PreviewID: "1",
  Config: &MarketOrder{ // or any other implementation
    BaseSize: "1",
    QuoteSize: "1",
  },
})
```

#### Cancel order(s)

```go
err := client.CancelOrders(ctx, "id", "id2") // returns any failures, any error
```

#### List orders

Partial implementation; no query parameters are allowed.
Implementing the websocket version is more important for now

```go
order, err := client.ListOrders(ctx)
```
