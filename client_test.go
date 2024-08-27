package coinbase

import (
	"crypto/ecdsa"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"
)

var mockPK *ecdsa.PrivateKey

func init() {
	var err error
	mockPK, err = parsePK(`-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIAh5qA3rmqQQuu0vbKV/+zouz/y/Iy2pLpIcWUSyImSwoAoGCCqGSM49
AwEHoUQDQgAEYD54V/vp+54P9DXarYqx4MPcm+HKRIQzNasYSoRQHQ/6S6Ps8tpM
cT+KvIIC8W/e9k0W7Cm72M1P9jU7SLf/vg==
-----END EC PRIVATE KEY-----`)

	if err != nil {
		panic(err)
	}

}

type testserver struct {
	c      *Client
	server *httptest.Server
}

func newTestserver(status int, mock any) *testserver {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf, err := json.Marshal(mock)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(status)
		w.Write(buf)
	}))

	return &testserver{
		c: &Client{
			baseURL:    s.URL,
			httpClient: &http.Client{Timeout: time.Second * 4},
			keyName:    "test",
			key:        mockPK,
		},
		server: s,
	}
}
