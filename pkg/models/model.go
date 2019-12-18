package models

type Model interface {
	Path() string
	Validate() bool
	GetData() map[string]interface{}
	GetName() string
}
