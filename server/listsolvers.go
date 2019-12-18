package server

import (
	"fmt"

	"context"

	"github.com/menta2l/lcm"
)

func (s LcmServer) Listsolvers(ctx context.Context, r *lcm.ListSolversRequest) (*lcm.ListSolversResponse, error) {
	fmt.Println("List solvers")
	keys, err := s.Backend.GetKeys("kv/lvm/solvers")
	if err != nil {
		return nil, err
	}
	var solvers []*lcm.SolverShort
	for _, key := range keys {
		res, _ := s.Backend.Read("kv/lvm/solvers/" + key)
		solvers = append(solvers, &lcm.SolverShort{Type: res.Data["data"].(map[string]interface{})["type"].(string), Name: res.Data["data"].(map[string]interface{})["name"].(string)})
	}
	return &lcm.ListSolversResponse{Solver: solvers}, nil
	//	return nil, errors.New("not yet implemented")
}
