package models

import (
	"errors"
	"fmt"

	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/cloudflare"
	"github.com/go-acme/lego/v3/providers/dns/gcloud"
	"github.com/menta2l/lcm"
	"github.com/menta2l/lcm/pkg/vault"
)

type Solver struct {
	BaseModel
	Email  string "json:email"
	Solver struct {
		ApiKey  string `json:"apiKey,omitempty"`
		Sa      string `json:"sa,omitempty"`
		Project string `json:"project,omitempty"`
	}
}

func NewSolver(item *lcm.Solver, client *vault.Client) (*Solver, error) {
	data := vault.StructToMap(item)
	switch item.GetDnsSolvers().(type) {
	case *lcm.Solver_CloudFlareSolver:
		data["type"] = "cloudflare"
		data["solver"] = vault.StructToMap(item.GetCloudFlareSolver())
	case *lcm.Solver_GoogleCloudSolver:
		data["solver"] = vault.StructToMap(item.GetGoogleCloudSolver())
		data["type"] = "gcloud"
	}
	_, ok := data["dns_solvers"]
	if ok {
		delete(data, "dns_solvers")
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
	s.Name = s.data["name"].(string)
	return true, nil
}
func (s *Solver) GetProvider() (challenge.Provider, error) {
	fmt.Printf("Solver type : %s %s \n", s.Type, s.Name)
	switch s.Type {
	case "cloudflare":
		cfg := cloudflare.NewDefaultConfig()
		cfg.AuthEmail = s.Email
		cfg.AuthKey = s.Solver.ApiKey
		return cloudflare.NewDNSProviderConfig(cfg)
	case "gcloud":
		return gcloud.NewDNSProviderServiceAccountKey([]byte(s.Solver.Sa))
	}
	return nil, fmt.Errorf("unknown provider")
}
func (s *Solver) Unmarshal(data map[string]interface{}) {
	s.unmarshal(data)
	if val, ok := data["email"]; ok {
		s.Email = val.(string)
	}

	if val, ok := data["solver"]; ok {
		if val, ok := val.(map[string]interface{})["email"]; ok {
			s.Email = val.(string)
		}
		if val, ok := val.(map[string]interface{})["apikey"]; ok {
			s.Solver.ApiKey = val.(string)
		}
		if val, ok := val.(map[string]interface{})["project"]; ok {
			s.Solver.Project = val.(string)
		}
		if val, ok := val.(map[string]interface{})["sa"]; ok {
			s.Solver.Sa = val.(string)
		}

	}

}
