package models

import (
	"errors"

	"github.com/menta2l/lcm/pkg/vault"
)

type BaseModel struct {
	data   map[string]interface{}
	path   string
	client *vault.Client
	Type   string `json:"type"`
	Name   string `json:"name"`
}

func (m *BaseModel) Save(update bool) error {
	if update == false {
		exist := m.client.Exist(m.Path())
		if exist {
			return errors.New("record already exist " + m.Path())
		}
	}
	return m.client.Write(m.Path(), m.data)
}
func (m *BaseModel) Path() string {
	return m.path + m.Name
}
func (m *BaseModel) unmarshal(data map[string]interface{}) {
	if val, ok := data["name"]; ok {
		m.Name = val.(string)
	}
	if val, ok := data["type"]; ok {
		m.Type = val.(string)
	}
	m.data = data
}
