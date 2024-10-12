package intf

import "src/module/config/schema"

type VariableRepo interface {
	ListVariable() ([]schema.Variable, error)
	RetrieveVariable(key int) (*schema.Variable, error)
	CreateVariable(variable *schema.Variable) (*schema.Variable, error)
	UpdateVariable(key int, variable map[string]interface{}) (*schema.Variable, error)
	DeleteVariable(key int) error
	DeleteListVariable(keys []int) error
}
