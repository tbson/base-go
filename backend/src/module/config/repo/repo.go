package repo

import (
	"src/module/config/schema"
	"src/util/dbutil"
	"src/util/errutil"
)

type VariableRepo struct{}

func (vr VariableRepo) CreateVariable(variable *schema.Variable) (*schema.Variable, error) {
	result := dbutil.Db().Create(variable)
	err := result.Error
	if err != nil {
		return variable, errutil.NewGormError(err)
	}
	return variable, err
}
