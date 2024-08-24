Go Coinbase Pro [![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/AnthonyHewins/coinbase)
========

v3 [CoinBase Pro](https://pro.coinbase.com) API client fork

- [Go Coinbase Pro ](#go-coinbase-pro-)
    - [Older Go versions](#older-go-versions)
  - [Documentation](#documentation)
    - [Decimal management](#decimal-management)
    - [Retry](#retry)
    - [Examples](#examples)
      - [Create order](#create-order)
      - [Cancel order(s)](#cancel-orders)
      - [List orders](#list-orders)
    - [Websockets](#websockets)

### Older Go versions
```sh
go get github.com/AnthonyHewins/coinbase
```

## Documentation
For full details on functionality, see [GoDoc](http://godoc.org/github.com/preichenberger/go-coinbasepro) documentation.

### Decimal management

To manage precision correctly, this library sends all price values as strings for now.
Considering building a decimal library directly into the codebase, but not sure which one

### Retry
You can set a retry count which uses exponential backoff: (2^(retry_attempt) - 1) / 2 * 1000 * milliseconds
```
client.RetryCount = 3 # 500ms, 1500ms, 3500ms
```

### Examples

#### Create order

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

### Websockets
Listen for websocket messages

```go
  import(
    ws "github.com/gorilla/websocket"
  )

  var wsDialer ws.Dialer
  wsConn, _, err := wsDialer.Dial("wss://ws-feed.pro.coinbase.com", nil)
  if err != nil {
    println(err.Error())
  }

  subscribe := coinbasepro.Message{
    Type:      "subscribe",
    Channels: []coinbasepro.MessageChannel{
      coinbasepro.MessageChannel{
        Name: "heartbeat",
        ProductIds: []string{
          "BTC-USD",
        },
      },
      coinbasepro.MessageChannel{
        Name: "level2",
        ProductIds: []string{
          "BTC-USD",
        },
      },
    },
  }
  if err := wsConn.WriteJSON(subscribe); err != nil {
    println(err.Error())
  }

  for true {
    message := coinbasepro.Message{}
    if err := wsConn.ReadJSON(&message); err != nil {
      println(err.Error())
      break
    }

    println(message.Type)
  }

```
