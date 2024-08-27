package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGettingAccounts(mainTest *testing.T) {
	c := testClient()

	t := assert.New(mainTest)

	accts, err := c.ListAccounts(context.Background())
	if !t.Nil(err, "should not error getting an account") {
		return
	}

	if !t.NotEmpty(accts, "user should have at least one account") {
		return
	}

	acct, err := c.GetAccount(context.Background(), accts[0].ID)
	if !t.Nil(err, "should not error getting account %s", accts[0].ID) {
		return
	}

	t.Equal(&accts[0], acct, "the account fetched during the list should match the fields of the regular GET, in terms of marshal/unmarshal")
}
