package coinbase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const DefaultProdURL = "https://api.coinbase.com/api/v3/brokerage"

// const DefaultSandboxURL = "https://api.sandbox.pro.coinbase.com"

type Client struct {
	baseURL    string
	httpClient *http.Client

	keyName, keySecret string
}

func NewClient(baseURL, keySecret, keyName string, httpClient *http.Client) (*Client, error) {
	switch {
	case len(keyName) == 0:
		return nil, fmt.Errorf("key name missing for coinbase client")
	case len(keySecret) == 0:
		return nil, fmt.Errorf("key secret not found for coinbase client")
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
		keyName:    keyName,
		keySecret:  keySecret,
	}, nil
}

func (c *Client) request(ctx context.Context, method string, url string, params, result interface{}) (res *http.Response, err error) {
	jwtStr, err := c.generateToken(method, url)
	if err != nil {
		return nil, err
	}

	var body io.Reader

	if params != nil {
		data, err := json.Marshal(params)
		if err != nil {
			return res, err
		}

		body = bytes.NewReader(data)
	}

	fullURL := fmt.Sprintf("%s%s", c.baseURL, url)
	req, err := http.NewRequestWithContext(ctx, method, fullURL, body)
	if err != nil {
		return res, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+jwtStr)

	res, err = c.httpClient.Do(req)
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
		if err = json.NewDecoder(res.Body).Decode(result); err != nil {
			return res, err
		}
	}

	return res, nil
}
