package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGettingAccounts(mainTest *testing.T) {
	c := testClient()

	t := assert.New(mainTest)

	accts, err := c.ListAccounts(context.Background())
	if !t.Nil(err, "should not error getting accounts") {
		return
	}

	if !t.NotEmpty(accts, "user should have at least one account") {
		return
	}

	acct, err := c.GetAccount(context.Background(), accts[0].ID)
	if !t.Nil(err, "should not error getting account %s", accts[0].ID) {
		return
	}

	accts[0].Updated = time.Time{}
	acct.Updated = time.Time{}
	accts[0].Created = time.Time{}
	acct.Created = time.Time{}

	t.Equal(&accts[0], acct, "the account fetched during the list should match the fields of the regular GET, in terms of marshal/unmarshal")
}
