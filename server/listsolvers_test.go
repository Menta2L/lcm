package server

import (
	"testing"

	"context"
	"github.com/menta2l/lcm"
	"github.com/stretchr/testify/assert"
)

func TestListsolvers(t *testing.T) {
	ctx := context.Background()
	req := &lcm.ListSolversRequest{}

	res, err := cli.Listsolvers(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
