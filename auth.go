package coinbase

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
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
	block, _ := pem.Decode([]byte(c.keySecret))
	if block == nil {
		return "", fmt.Errorf("jwt: Could not decode private key")
	}

	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("jwt: %w", err)
	}

	sig, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.ES256, Key: key},
		(&jose.SignerOptions{NonceSource: nonceSource{}}).WithType("JWT").WithHeader("kid", c.keyName),
	)
	if err != nil {
		return "", fmt.Errorf("jwt: %w", err)
	}

	cl := &claims{
		Claims: &jwt.Claims{
			Subject:   c.keyName,
			Issuer:    "cdp",
			NotBefore: jwt.NewNumericDate(time.Now()),
			Expiry:    jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
		},
		URI: fmt.Sprintf("%s %s%s", method, hostSign, path),
	}
	jwtString, err := jwt.Signed(sig).Claims(cl).CompactSerialize()
	if err != nil {
		return "", fmt.Errorf("jwt: %w", err)
	}
	return jwtString, nil
}
