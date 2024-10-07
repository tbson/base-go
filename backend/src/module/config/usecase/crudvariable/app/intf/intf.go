package intf

import "src/module/config/schema"

type VariableRepo interface {
	CreateVariable(variable *schema.Variable) (*schema.Variable, error)
}
