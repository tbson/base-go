package repo

import (
	"src/common/ctype"
	"src/module/config/schema"
	"src/util/errutil"

	"gorm.io/gorm"
)

type VariableRepo struct {
	db *gorm.DB
}

func (r VariableRepo) New(db *gorm.DB) VariableRepo {
	return VariableRepo{db: db}
}

func (r VariableRepo) ListVariable(params ctype.Dict) ([]schema.Variable, error) {
	db := r.db.Order("id DESC")
	var variables []schema.Variable

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

func (r VariableRepo) RetrieveVariable(id int) (*schema.Variable, error) {
	var variable schema.Variable
	result := r.db.Where("id = ?", id).First(&variable)
	err := result.Error
	if err != nil {
		return &variable, errutil.NewGormError(err)
	}
	return &variable, err
}

func (r VariableRepo) CreateVariable(variable *schema.Variable) (*schema.Variable, error) {
	result := r.db.Create(variable)
	err := result.Error
	if err != nil {
		return variable, errutil.NewGormError(err)
	}
	return variable, err
}

func (r VariableRepo) UpdateVariable(id int, variable ctype.Dict) (*schema.Variable, error) {
	item, err := r.RetrieveVariable(id)
	if err != nil {
		return nil, err
	}
	result := r.db.Model(&item).Updates(variable)
	err = result.Error
	if err != nil {
		return nil, errutil.NewGormError(err)
	}
	return item, err
}

func (r VariableRepo) DeleteVariable(id int) ([]int, error) {
	ids := []int{id}
	result := r.db.Where("id = ?", id).Delete(&schema.Variable{})
	err := result.Error
	if err != nil {
		return ids, errutil.NewGormError(err)
	}
	return ids, err
}

func (r VariableRepo) DeleteListVariable(ids []int) ([]int, error) {
	result := r.db.Where("id IN (?)", ids).Delete(&schema.Variable{})
	err := result.Error
	if err != nil {
		return ids, errutil.NewGormError(err)
	}
	return ids, err
}
