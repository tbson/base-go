package variable

import (
	"src/common/ctype"
	"src/module/config/schema"
	"src/util/errutil"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func (r Repo) New(db *gorm.DB) Repo {
	return Repo{db: db}
}

func (r Repo) List(params ctype.Dict) ([]schema.Variable, error) {
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

func (r Repo) Retrieve(id int) (*schema.Variable, error) {
	var variable schema.Variable
	result := r.db.Where("id = ?", id).First(&variable)
	err := result.Error
	if err != nil {
		return &variable, errutil.NewGormError(err)
	}
	return &variable, err
}

func (r Repo) Create(variable *schema.Variable) (*schema.Variable, error) {
	result := r.db.Create(variable)
	err := result.Error
	if err != nil {
		return variable, errutil.NewGormError(err)
	}
	return variable, err
}

func (r Repo) Update(id int, variable ctype.Dict) (*schema.Variable, error) {
	item, err := r.Retrieve(id)
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

func (r Repo) Delete(id int) ([]int, error) {
	ids := []int{id}
	result := r.db.Where("id = ?", id).Delete(&schema.Variable{})
	err := result.Error
	if err != nil {
		return ids, errutil.NewGormError(err)
	}
	return ids, err
}

func (r Repo) DeleteList(ids []int) ([]int, error) {
	result := r.db.Where("id IN (?)", ids).Delete(&schema.Variable{})
	err := result.Error
	if err != nil {
		return ids, errutil.NewGormError(err)
	}
	return ids, err
}
