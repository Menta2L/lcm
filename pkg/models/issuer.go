package models

import (
	"errors"
	"fmt"

	"github.com/menta2l/lcm"
	"github.com/menta2l/lcm/pkg/vault"
)

type Issuer struct {
	BaseModel
	Issuer struct {
		Email     string "json:email"
		Server    string "json:server"
		SolverRef string "json:solverRef"
	}
}

func NewIssuer(item *lcm.IssuerRequest, client *vault.Client) (*Issuer, error) {
	data := vault.StructToMap(item)
	switch item.GetIssuer().(type) {
	case *lcm.IssuerRequest_LetsEncryptIssuer:
		data["issuer"] = vault.StructToMap(item.GetLetsEncryptIssuer())
		data["type"] = "lets-encrypt"
	case *lcm.IssuerRequest_SelfSignedIssuer:
		data["type"] = "self-signed"
		data["issuer"] = vault.StructToMap(item.GetSelfSignedIssuer())

	}

	item.GetIssuer()
	issuer := &Issuer{
		BaseModel: BaseModel{
			data:   data,
			client: client,
			path:   "kv/lcm/issuers/",
		},
	}

	_, err := issuer.Validate()
	if err != nil {
		return nil, err
	}
	return issuer, nil

}
func (i *Issuer) Validate() (bool, error) {
	if _, ok := i.data["name"]; !ok {
		return false, errors.New("name parameter is required")
	}
	i.Name = i.data["name"].(string)
	if i.data["type"] != "self-signed" {
		if i.client.Exist("kv/lcm/solvers/"+i.GetSolverRef()) == false {
			return false, fmt.Errorf("solver %s not exist", i.GetSolverRef())
		}
	}

	return true, nil
}
func (i *Issuer) GetData() map[string]interface{} {
	return i.data
}
func (i *Issuer) GetName() string {
	return i.data["name"].(string)
}
func (i *Issuer) GetSlover() (*Solver, error) {
	secret, err := i.client.Read("kv/lcm/solvers/" + i.GetSolverRef())
	if err != nil {
		return nil, err
	}
	//	jsonObj, err := json.Marshal(secret.Data["data"])
	solver := &Solver{}
	solver.Unmarshal(secret.Data["data"].(map[string]interface{}))
	//	err = json.Unmarshal(jsonObj, &solver)
	//	err = mapstructure.Decode(secret.Data["data"], &solver)
	//	if err != nil {
	//		return nil, err
	//	}

	return solver, nil
}
func (i *Issuer) GetSolverRef() string {
	return i.data["issuer"].(map[string]interface{})["solverRef"].(string)
}
func (m *Issuer) Unmarshal(data map[string]interface{}) {
	m.unmarshal(data)
	if val, ok := data["issuer"]; ok {
		if val, ok := val.(map[string]interface{})["email"]; ok {
			m.Issuer.Email = val.(string)
		}
		if val, ok := val.(map[string]interface{})["server"]; ok {
			m.Issuer.Server = val.(string)
		}
		if val, ok := val.(map[string]interface{})["solverRef"]; ok {
			m.Issuer.SolverRef = val.(string)
		}

	}

}
