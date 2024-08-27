package coinbase

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"time"

	"gopkg.in/go-jose/go-jose.v2"
	"gopkg.in/go-jose/go-jose.v2/jwt"
)

// the exact string coinbase wants to see in the URI claims
// for some reason
const hostSign = "api.coinbase.com/api/v3/brokerage"

var maxRand = big.NewInt(math.MaxInt64)

type claims struct {
	*jwt.Claims
	URI string `json:"uri"`
}

type nonceSource struct{}

func (n nonceSource) Nonce() (string, error) {
	r, err := rand.Int(rand.Reader, maxRand)
	if err != nil {
		return "", err
	}
	return r.String(), nil
}

// generate a one-time use JWT. You actually need to create a new JWT per request;
// this requires using the private key to sign a new token repeatedly
func (c *Client) generateToken(method, path string) (string, error) {
	sig, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.ES256, Key: c.key},
		(&jose.SignerOptions{NonceSource: nonceSource{}}).WithType("JWT").WithHeader("kid", c.keyName),
	)
	if err != nil {
		return "", fmt.Errorf("jwt: %w", err)
	}

	cl := &claims{
		Claims: &jwt.Claims{
			Issuer:    "cdp",
			Subject:   c.keyName,
			Expiry:    jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		URI: fmt.Sprintf("%s %s%s", method, hostSign, path),
	}

	return jwt.Signed(sig).Claims(cl).CompactSerialize()
}
