package coinbase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"time"
)

const DefaultProdURL = "https://api.coinbase.com/api/v3/brokerage"

// const DefaultSandboxURL = "https://api-public.sandbox.pro.coinbase.com"

type Client struct {
	BaseURL    string
	Secret     string
	Key        string
	Passphrase string
	HTTPClient *http.Client
	RetryCount int
}

type ClientConfig struct {
	BaseURL    string
	Key        string
	Passphrase string
	Secret     string
}

func NewClient(baseURL, key, passphrase, secret string, httpClient *http.Client) *Client {
	return &Client{
		BaseURL:    baseURL,
		Key:        key,
		Passphrase: passphrase,
		Secret:     secret,
		HTTPClient: httpClient,
		RetryCount: 0,
	}
}

func (c *Client) Request(ctx context.Context, method string, url string, params, result interface{}) (res *http.Response, err error) {
	for i := 0; i < c.RetryCount+1; i++ {
		retryDuration := time.Duration((math.Pow(2, float64(i))-1)/2*1000) * time.Millisecond
		time.Sleep(retryDuration)

		res, err = c.request(ctx, method, url, params, result)
		if res != nil && res.StatusCode == 429 {
			continue
		} else {
			break
		}
	}

	return res, err
}

func (c *Client) request(ctx context.Context, method string, url string,
	params, result interface{}) (res *http.Response, err error) {
	var data []byte
	var body io.Reader

	if params != nil {
		data, err = json.Marshal(params)
		if err != nil {
			return res, err
		}

		body = bytes.NewReader(data)
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, url)
	req, err := http.NewRequest(method, fullURL, body)
	if err != nil {
		return res, err
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Go Coinbase Pro Client 1.0")

	h, err := c.signature(method, url, timestamp, string(data))
	if err != nil {
		return res, err
	}

	for k, v := range h {
		req.Header.Add(k, v)
	}

	res, err = c.HTTPClient.Do(req)
	if err != nil {
		return res, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		coinbaseError := Error{}
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&coinbaseError); err != nil {
			return res, err
		}

		return res, coinbaseError
	}

	if result != nil {
		decoder := json.NewDecoder(res.Body)
		if err = decoder.Decode(result); err != nil {
			return res, err
		}
	}

	return res, nil
}

func (c *Client) signature(method, url, timestamp, data string) (map[string]string, error) {
	h := make(map[string]string)
	h["CB-ACCESS-KEY"] = c.Key
	h["CB-ACCESS-PASSPHRASE"] = c.Passphrase
	h["CB-ACCESS-TIMESTAMP"] = timestamp

	message := fmt.Sprintf(
		"%s%s%s%s",
		timestamp,
		method,
		url,
		data,
	)

	sig, err := generateSig(message, c.Secret)
	if err != nil {
		return nil, err
	}
	h["CB-ACCESS-SIGN"] = sig
	return h, nil
}
