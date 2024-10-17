package repo

import (
	"src/common/ctype"
	"src/module/config/schema"
	"src/util/dbutil"
	"src/util/errutil"
)

type VariableRepo struct{}

func (vr VariableRepo) ListVariable(params ctype.Dict) ([]schema.Variable, error) {
	var variables []schema.Variable
	db := dbutil.Db().Order("id DESC")

	if len(params) > 0 {
		db = db.Where(params)
	}
	result := db.Find(&variables)
	err := result.Error
	if err != nil {
		return variables, errutil.NewGormError(err)
	}
	return variables, err
}

func (vr VariableRepo) RetrieveVariable(id int) (*schema.Variable, error) {
	var variable schema.Variable
	result := dbutil.Db().Where("id = ?", id).First(&variable)
	err := result.Error
	if err != nil {
		return &variable, errutil.NewGormError(err)
	}
	return &variable, err
}

func (vr VariableRepo) CreateVariable(variable *schema.Variable) (*schema.Variable, error) {
	result := dbutil.Db().Create(variable)
	err := result.Error
	if err != nil {
		return variable, errutil.NewGormError(err)
	}
	return variable, err
}

func (vr VariableRepo) UpdateVariable(id int, variable ctype.Dict) (*schema.Variable, error) {
	item, err := vr.RetrieveVariable(id)
	if err != nil {
		return nil, err
	}
	result := dbutil.Db().Model(&item).Updates(variable)
	err = result.Error
	if err != nil {
		return nil, errutil.NewGormError(err)
	}
	return item, err
}

func (vr VariableRepo) DeleteVariable(id int) ([]int, error) {
	ids := []int{id}
	result := dbutil.Db().Where("id = ?", id).Delete(&schema.Variable{})
	err := result.Error
	if err != nil {
		return ids, errutil.NewGormError(err)
	}
	return ids, err
}

func (vr VariableRepo) DeleteListVariable(ids []int) ([]int, error) {
	result := dbutil.Db().Where("id IN (?)", ids).Delete(&schema.Variable{})
	err := result.Error
	if err != nil {
		return ids, errutil.NewGormError(err)
	}
	return ids, err
}
