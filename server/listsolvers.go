package server

import (
	"errors"

	"context"
	"github.com/menta2l/lcm"
)

func (s LcmServer) Listsolvers(ctx context.Context, r *lcm.ListSolversRequest) (*lcm.ListSolversResponse, error) {
	return nil, errors.New("not yet implemented")
}
