package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPortfolio(mainTest *testing.T) {
	c := testClient()

	t := assert.New(mainTest)
	p, err := c.ListPortfolios(context.Background())

	if !t.NoError(err, "should not fail listing portfolios") {
		return
	}

	if !t.NotEmpty(p, "should have at least one portfolio") {
		return
	}

	_, err = c.GetPortfolio(context.Background(), p[0].UUID)
	if !t.NoError(err, "should not error getting portfolio %s", p[0].UUID) {
		return
	}
}
