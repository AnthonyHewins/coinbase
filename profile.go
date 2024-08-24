package coinbase

import (
	"context"
	"fmt"
)

type Profile struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	Active    bool   `json:"active"`
	IsDefault bool   `json:"is_default"`
	CreatedAt Time   `json:"created_at,string"`
}

type ProfileTransfer struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

// Client Funcs

// GetProfiles retrieves a list of profiles
func (c *Client) GetProfiles(ctx context.Context) ([]Profile, error) {
	var profiles []Profile

	url := fmt.Sprintf("/profiles")
	_, err := c.Request(ctx, "GET", url, nil, &profiles)
	return profiles, err
}

// GetProfile retrieves a single profile
func (c *Client) GetProfile(ctx context.Context, id string) (Profile, error) {
	var profile Profile

	url := fmt.Sprintf("/profiles/%s", id)
	_, err := c.Request(ctx, "GET", url, nil, &profile)
	return profile, err
}

// CreateProfileTransfer transfers a currency amount from one profile to another
func (c *Client) CreateProfileTransfer(ctx context.Context, newTransfer *ProfileTransfer) error {
	url := fmt.Sprintf("/profiles/transfer")
	_, err := c.Request(ctx, "POST", url, newTransfer, nil)

	return err
}
