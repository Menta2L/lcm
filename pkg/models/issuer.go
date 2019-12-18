package models

import (
	"errors"
	"fmt"

	"github.com/menta2l/lcm"
	"github.com/menta2l/lcm/pkg/vault"
)

type Issuer struct {
	BaseModel
	solverRef string
}

func NewIssuer(item *lcm.IssuerRequest, client *vault.Client) (*Issuer, error) {
	data := vault.StructToMap(item)
	solverRef := ""
	switch item.GetIssuer().(type) {
	case *lcm.IssuerRequest_LetsEncryptIssuer:
		solverRef = item.GetLetsEncryptIssuer().GetSolverRef()
		data["type"] = "lets-encrypt"
	case *lcm.IssuerRequest_SelfSignedIssuer:
		data["type"] = "self-signed"
	}
	item.GetIssuer()
	issuer := &Issuer{
		BaseModel: BaseModel{
			data:   data,
			client: client,
			path:   "kv/lcm/issuers/",
		},
		solverRef: solverRef,
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
	i.name = i.data["name"].(string)
	if i.data["type"] != "self-signed" {
		if i.solverRef == "" {
			return false, errors.New("solver  parameter is required")
		}
		if i.client.Exist("kv/lcm/solvers/"+i.solverRef) == false {
			return false, fmt.Errorf("solver %s not exist", i.solverRef)
		}
	}
	i.data["solverRef"] = i.solverRef

	return true, nil
}
func (i *Issuer) GetData() map[string]interface{} {
	return i.data
}
func (i *Issuer) GetName() string {
	return i.data["name"].(string)
}
