package repo

import (
	"src/module/config/schema"
	"src/util/dbutil"
)

type VariableRepo struct{}

func (vr VariableRepo) CreateVariable(variable *schema.Variable) (*schema.Variable, error) {
	result := dbutil.Db().Create(variable)
	return variable, result.Error
}
