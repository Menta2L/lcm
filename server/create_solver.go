package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/menta2l/lcm"
	"github.com/menta2l/lcm/pkg/models"
)

func (s LcmServer) CreateSolver(ctx context.Context, r *lcm.Solver) (*empty.Empty, error) {
	solver, err := models.NewSolver(r, s.Backend)
	if err != nil {
		return nil, err
	}
	err = solver.Save(false)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
