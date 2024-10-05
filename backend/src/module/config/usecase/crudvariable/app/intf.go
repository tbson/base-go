package core

import "src/module/config/repo/schema"

type VariableRepo interface {
	CreateVariable(variable *schema.Variable) (*schema.Variable, error)
}
