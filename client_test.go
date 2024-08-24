package coinbase

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"
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

	return &testserver{
		c: &Client{
			BaseURL:    s.URL,
			HTTPClient: &http.Client{Timeout: time.Second * 4},
		},
		server: s,
	}
}
