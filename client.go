package coinbase

import (
	"bytes"
	"context"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
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
	baseURL    string
	secret     string
	key        string
	passphrase string
	httpClient *http.Client
	retryCount int

	tokenCache tokenCache
}

func NewClient(baseURL, key, passphrase, secret string, httpClient *http.Client) (*Client, error) {
	block, _ := pem.Decode([]byte(keySecret))
	if block == nil {
		return "", fmt.Errorf("jwt: Could not decode private key")
	}

	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("jwt: %w", err)
	}

	sig, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.ES256, Key: key},
		(&jose.SignerOptions{NonceSource: nonceSource{}}).WithType("JWT").WithHeader("kid", keyName),
	)
	if err != nil {
		return "", fmt.Errorf("jwt: %w", err)
	}

	cl := &APIKeyClaims{
		Claims: &jwt.Claims{
			Subject:   keyName,
			Issuer:    "cdp",
			NotBefore: jwt.NewNumericDate(time.Now()),
			Expiry:    jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
		},
		URI: uri,
	}
	jwtString, err := jwt.Signed(sig).Claims(cl).CompactSerialize()
	if err != nil {
		return "", fmt.Errorf("jwt: %w", err)
	}

	return &Client{
		baseURL:    baseURL,
		key:        key,
		passphrase: passphrase,
		secret:     secret,
		httpClient: httpClient,
		retryCount: 0,
	}, nil
}

// If you have a cached token you've saved and you want to avoid re-authenticating,
// use this
func (c *Client) SetToken(s string) {
	c.tokenCache.setToken(s)
}

// Get the current token that's cached by the client, if there is one,
// in the event you want to cache the token outside this library
func (c *Client) GetCachedToken() string {
	return c.tokenCache.getToken()
}

func (c *Client) Request(ctx context.Context, method string, url string, params, result interface{}) (res *http.Response, err error) {
	for i := 0; i < c.retryCount+1; i++ {
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

	fullURL := fmt.Sprintf("%s%s", c.baseURL, url)
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
		decoder := json.NewDecoder(res.Body)
		if err = decoder.Decode(result); err != nil {
			return res, err
		}
	}

	return res, nil
}

func (c *Client) signature(method, url, timestamp, data string) (map[string]string, error) {
	h := make(map[string]string)
	h["CB-ACCESS-KEY"] = c.key
	h["CB-ACCESS-PASSPHRASE"] = c.passphrase
	h["CB-ACCESS-TIMESTAMP"] = timestamp

	message := fmt.Sprintf(
		"%s%s%s%s",
		timestamp,
		method,
		url,
		data,
	)

	sig, err := generateSig(message, c.secret)
	if err != nil {
		return nil, err
	}
	h["CB-ACCESS-SIGN"] = sig
	return h, nil
}
