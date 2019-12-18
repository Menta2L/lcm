package models

import (
	"errors"
	"fmt"

	"github.com/menta2l/lcm/pkg/vault"
)

type BaseModel struct {
	name   string
	data   map[string]interface{}
	path   string
	client *vault.Client
}

func (m *BaseModel) Save(update bool) error {
	if update == false {
		exist := m.client.Exist(m.Path())
		if exist {
			return errors.New("record already exist " + m.Path())
		}
	}
	fmt.Println("Save")
	return m.client.Write(m.Path(), m.data)
}
func (m *BaseModel) Path() string {
	return m.path + m.name
}
