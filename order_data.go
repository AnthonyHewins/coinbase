package coinbase

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

//go:generate enumer -type StopDirection -transform snake-upper -json
type StopDirection byte

const (
	StopDirectionUnspecified StopDirection = iota
	StopDirectionStopUp                    // buy side
	StopDirectionStopDown                  // sell side
)

// wrapper type to encapsulate JSON fields the end developer doesn't
// need to worry about so the API is very clean
type orderData struct {
	// The amount of the second Asset in the Trading Pair. For example, on the
	// BTC/USD Order Book, USD is the Quote Asset.
	QuoteSize decimal.Decimal `json:"quote_size,omitempty"`

	// The amount of the first Asset in the Trading Pair. For example, on the
	// BTC-USD Order Book, BTC is the Base Asset.
	BaseSize decimal.Decimal `json:"base_size,omitempty"`

	// The specified price, or better, that the Order should be executed at. A
	// Buy Order will execute at or lower than the limit price. A Sell Order
	// will execute at or higher than the limit price.
	LimitPrice       decimal.Decimal `json:"limit_price,omitempty"`
	PostOnly         bool            `json:"post_only,omitempty"`
	End              *time.Time      `json:"end_time,omitempty"`
	Stop             decimal.Decimal `json:"stop_price,omitempty"`
	StopDirection    StopDirection   `json:"stop_direction,omitempty"`
	StopTriggerPrice decimal.Decimal `json:"stop_trigger_price,omitempty"`
}

func (o *orderData) MarshalJSON() ([]byte, error) {
	m := map[string]any{}

	type pairs struct {
		name  string
		value decimal.Decimal
	}

	for _, v := range []pairs{
		{"quote_size", o.QuoteSize},
		{"base_size", o.BaseSize},
		{"limit_price", o.LimitPrice},
		{"stop_price", o.Stop},
		{"stop_trigger_price", o.StopTriggerPrice},
	} {
		if !v.value.IsZero() {
			m[v.name] = v.value
		}
	}

	if o.PostOnly {
		m["post_only"] = true
	}

	if o.End != nil {
		m["end_time"] = *o.End
	}

	if o.StopDirection != StopDirectionUnspecified {
		m["stop_direction"] = o.StopDirection
	}

	return json.Marshal(m)
}

func (o *orderData) stopSide() Side {
	switch x := o.StopDirection; x {
	case StopDirectionStopUp:
		return SideBuy
	case StopDirectionStopDown:
		return SideSell
	default:
		return SideUnspecified
	}
}

func (o *orderData) end() time.Time {
	if o.End == nil {
		return time.Time{}
	}

	return *o.End
}

func discoverConfig(orderType string, b json.RawMessage) (OrderConfig, error) {
	if orderType == "MARKET" {
		conf := MarketOrder{}
		if err := json.Unmarshal(b, &conf); err != nil {
			return nil, fmt.Errorf("failed unmarshal of market order: %w", err)
		}

		return &conf, nil
	}

	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, fmt.Errorf("invalid order config, expected object (%w): %s", err, b)
	}

	var data *orderData
	switch orderType {
	case "LIMIT":
		for k, v := range m {
			if err := json.Unmarshal(v, &data); err != nil {
				return nil, fmt.Errorf("failed unmarshal of limit order (type %s): %w", k, err)
			}

			switch k {
			case "sor_limit_ioc":
				return &LimitOrderIOC{BaseSize: data.BaseSize, LimitPrice: data.LimitPrice}, nil
			case "limit_limit_gtc":
				return &LimitOrderGTC{
					BaseSize:   data.BaseSize,
					LimitPrice: data.LimitPrice,
					PostOnly:   data.PostOnly,
				}, nil
			case "limit_limit_gtd":
				return &LimitOrderGTD{
					BaseSize:   data.BaseSize,
					LimitPrice: data.LimitPrice,
					EndTime:    data.end(),
					PostOnly:   data.PostOnly,
				}, nil
			case "limit_limit_fok":
				return &LimitOrderFOK{BaseSize: data.BaseSize, LimitPrice: data.LimitPrice}, nil
			default:
				return nil, fmt.Errorf("invalid order configuration key: %s", k)
			}
		}
	case "STOP", "STOP_LIMIT":
		for k, v := range m {
			if err := json.Unmarshal(v, &data); err != nil {
				return nil, fmt.Errorf("failed unmarshal of stop limit order (type %s): %w", k, err)
			}

			switch side := data.stopSide(); k {
			case "stop_limit_stop_limit_gtc":
				return &StopLimitOrderGTC{
					BaseSize:   data.BaseSize,
					LimitPrice: data.LimitPrice,
					Stop:       data.Stop,
					Side:       side,
				}, nil
			case "stop_limit_stop_limit_gtd":
				return &StopLimitOrderGTD{
					BaseSize:   data.BaseSize,
					LimitPrice: data.LimitPrice,
					Stop:       data.Stop,
					EndTime:    data.end(),
					Side:       side,
				}, nil
			default:
				return nil, fmt.Errorf("invalid order configuration key: %s", k)
			}
		}
	case "BRACKET":
		for k, v := range m {
			if err := json.Unmarshal(v, &data); err != nil {
				return nil, fmt.Errorf("failed unmarshal of stop limit order (type %s): %w", k, err)
			}

			switch k {
			case "trigger_bracket_gtc":
				return &TriggerBracketOrderGTC{
					BaseSize:         data.BaseSize,
					LimitPrice:       data.LimitPrice,
					StopTriggerPrice: data.StopTriggerPrice,
				}, nil
			case "trigger_bracket_gtd":
				return &TriggerBracketOrderGTD{
					BaseSize:         data.BaseSize,
					LimitPrice:       data.LimitPrice,
					EndTime:          data.end(),
					StopTriggerPrice: data.StopTriggerPrice,
				}, nil
			default:
				return nil, fmt.Errorf("invalid bracket type %s", k)
			}
		}
	default:
		return nil, fmt.Errorf("unknown order type: %s", orderType)
	}

	return nil, fmt.Errorf("order config missing")
}
