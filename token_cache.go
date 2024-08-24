package coinbase

import "sync"

type tokenCache struct {
	mu    sync.RWMutex
	token string
}

func (t *tokenCache) setToken(s string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.token = s
}

func (t *tokenCache) getToken() string {
	t.mu.RLock()
	defer t.mu.Unlock()
	return t.token
}
