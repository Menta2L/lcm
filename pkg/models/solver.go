package models

import (
	"errors"

	"github.com/menta2l/lcm"
	"github.com/menta2l/lcm/pkg/vault"
)

type Solver struct {
	BaseModel
}

func NewSolver(item *lcm.Solver, client *vault.Client) (*Solver, error) {
	data := vault.StructToMap(item)
	switch item.GetDnsSolvers().(type) {
	case *lcm.Solver_CloudFlareSolver:
		data["type"] = "cloudflare"
	case *lcm.Solver_GoogleCloudSolver:
		data["type"] = "google"
	}
	solver := &Solver{
		BaseModel: BaseModel{
			data:   data,
			client: client,
			path:   "kv/lcm/solvers/",
		},
	}

	_, err := solver.Validate()
	if err != nil {
		return nil, err
	}
	return solver, nil

}
func (s *Solver) Validate() (bool, error) {
	if _, ok := s.data["name"]; !ok {
		return false, errors.New("name parameter is required")
	}
	s.name = s.data["name"].(string)
	return true, nil
}
