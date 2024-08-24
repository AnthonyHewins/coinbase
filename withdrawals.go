package coinbase

import (
	"context"
	"fmt"
)

type WithdrawalCrypto struct {
	Currency      string `json:"currency"`
	Amount        string `json:"amount"`
	CryptoAddress string `json:"crypto_address"`
	ProfileID     string `json:"profile_id"`
}

type WithdrawalCoinbase struct {
	Currency          string `json:"currency"`
	Amount            string `json:"amount"`
	CoinbaseAccountID string `json:"coinbase_account_id"`
}

func (c *Client) CreateWithdrawalCrypto(ctx context.Context, newWithdrawalCrypto *WithdrawalCrypto) (WithdrawalCrypto, error) {
	var savedWithdrawal WithdrawalCrypto
	url := fmt.Sprintf("/withdrawals/crypto")
	_, err := c.request(ctx, "POST", url, newWithdrawalCrypto, &savedWithdrawal)
	return savedWithdrawal, err
}

func (c *Client) CreateWithdrawalCoinbase(ctx context.Context, newWithdrawalCoinbase *WithdrawalCoinbase) (WithdrawalCoinbase, error) {
	var savedWithdrawal WithdrawalCoinbase
	url := fmt.Sprintf("/withdrawals/coinbase-account")
	_, err := c.request(ctx, "POST", url, newWithdrawalCoinbase, &savedWithdrawal)
	return savedWithdrawal, err
}
