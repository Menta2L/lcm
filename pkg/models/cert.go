package models

import (
	"errors"
	"fmt"

	"github.com/menta2l/lcm"
	"github.com/menta2l/lcm/pkg/vault"
)

type Cert struct {
	BaseModel
	issuerRef string
}

func NewCert(item *lcm.CertificateRequest, client *vault.Client) (*Cert, error) {
	data := vault.StructToMap(item)
	cert := &Cert{
		BaseModel: BaseModel{
			data:   data,
			client: client,
			path:   "kv/lcm/cert/",
		},
		issuerRef: item.GetIssuerRef(),
	}

	_, err := cert.Validate()
	if err != nil {
		return nil, err
	}
	return cert, nil

}
func (c *Cert) Validate() (bool, error) {
	if _, ok := c.data["name"]; !ok {
		return false, errors.New("name parameter is required")
	}
	c.name = c.data["name"].(string)
	if c.issuerRef == "" {
		return false, errors.New("issuerRef  parameter is required")
	}
	if c.client.Exist("kv/lcm/issuers/"+c.issuerRef) == false {
		return false, fmt.Errorf("solver %s not exist", c.issuerRef)
	}

	return true, nil
}
