package coinbase

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"
)

const (
	mockKey = "organizations/{org_id}/apiKeys/{key_id}"
	mockPK  = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIAh5qA3rmqQQuu0vbKV/+zouz/y/Iy2pLpIcWUSyImSwoAoGCCqGSM49
AwEHoUQDQgAEYD54V/vp+54P9DXarYqx4MPcm+HKRIQzNasYSoRQHQ/6S6Ps8tpM
cT+KvIIC8W/e9k0W7Cm72M1P9jU7SLf/vg==
-----END EC PRIVATE KEY-----`
)

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

	c, err := NewClient(s.URL, mockKey, mockPK, &http.Client{Timeout: time.Second})
	if err != nil {
		panic(err)
	}

	return &testserver{c: c, server: s}
}
