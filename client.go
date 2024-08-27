package coinbase

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const DefaultProdURL = "https://api.coinbase.com/api/v3/brokerage"

const DefaultSandboxURL = "https://api.sandbox.coinbase.com"

type Client struct {
	baseURL    string
	httpClient *http.Client

	keyName string
	key     *ecdsa.PrivateKey
}

func parsePK(keySecret string) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(keySecret))
	if block == nil {
		return nil, fmt.Errorf("could not decode coinbase's ECDSA private key (key length: %d)", len(keySecret))
	}

	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed parsing coinbase's x509 EC private key: %w", err)
	}

	return key, nil
}

func NewClient(baseURL, keyName, keySecret string, httpClient *http.Client) (*Client, error) {
	key, err := parsePK(keySecret)
	if err != nil {
		return nil, err
	}
	return NewClientWithPrivateKey(baseURL, keyName, key, httpClient)
}

func NewClientWithPrivateKey(baseURL, keyName string, privateKey *ecdsa.PrivateKey, httpClient *http.Client) (*Client, error) {
	switch {
	case len(keyName) == 0:
		return nil, fmt.Errorf("key name missing for coinbase client: empty string")
	case privateKey == nil:
		return nil, fmt.Errorf("private key missing for coinbase client")
	}

	keynameParts := strings.Split(keyName, "/")
	if len(keynameParts) != 4 || keynameParts[0] != "organizations" || keynameParts[2] != "apiKeys" {
		return nil, fmt.Errorf(
			"keyname supplied is invalid: must follow the format organizations/{keyname}/apiKeys/{keyid}",
		)
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
		keyName:    keyName,
		key:        privateKey,
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

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 401:
		return nil, fmt.Errorf("request was unauthorized (HTTP 401)")
	case 200:
		if result != nil {
			err = json.Unmarshal(buf, result)
		}

		return res, err
	}

	if res.StatusCode != 200 {
		coinbaseError := Error{}
		if err := json.Unmarshal(buf, &coinbaseError); err != nil {
			return res, &UnmarshalErr{
				Err:      err,
				Buf:      string(buf),
				RespCode: res.StatusCode,
			}
		}

		return res, coinbaseError
	}

	return res, nil
}
