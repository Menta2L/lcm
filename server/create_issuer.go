package server

import (
	"errors"

	"context"
	"github.com/menta2l/lcm"
)

func (s LcmServer) CreateIssuer(ctx context.Context, r *lcm.IssuerRequest) (*lcm.IssuerResponse, error) {
	return nil, errors.New("not yet implemented")
}
