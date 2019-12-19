package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/menta2l/lcm"
	"github.com/menta2l/lcm/pkg/models"
)

func (s LcmServer) CreateIssuer(ctx context.Context, r *lcm.IssuerRequest) (*empty.Empty, error) {
	iss, err := models.NewIssuer(r, s.Backend)
	if err != nil {
		return nil, err
	}
	err = iss.Save(false)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
