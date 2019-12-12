package server

import (
	"testing"

	"context"
	"github.com/menta2l/lcm"
	"github.com/stretchr/testify/assert"
)

func TestCreateSolver(t *testing.T) {
	ctx := context.Background()
	req := &lcm.CreateSolverRequest{}

	res, err := cli.CreateSolver(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
