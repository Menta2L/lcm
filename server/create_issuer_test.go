package server

import (
	"testing"

	"context"
	"github.com/menta2l/lcm"
	"github.com/stretchr/testify/assert"
)

func TestCreateIssuer(t *testing.T) {
	ctx := context.Background()
	req := &lcm.IssuerRequest{}

	res, err := cli.CreateIssuer(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
