package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListPortfolios(mainTest *testing.T) {
	c := testClient()

	t := assert.New(mainTest)
	p, err := c.ListPortfolios(context.Background())

	if !t.NoError(err, "should not fail listing portfolios") {
		return
	}

	if !t.NotEmpty(p, "should have at least one portfolio") {
		return
	}
}

func TestGetPortfolio(mainTest *testing.T) {
	c := testClient()

	t := assert.New(mainTest)

	_, err := c.GetPortfolio(context.Background(), singleton.Portfolio)
	if !t.NoError(err, "should not error getting portfolio %s", singleton.Portfolio) {
		return
	}
}
