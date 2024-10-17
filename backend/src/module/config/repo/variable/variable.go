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

type Variable = schema.Variable

func (r Repo) New(db *gorm.DB) Repo {
	return Repo{db: db}
}

func (r Repo) List(params ctype.Dict) ([]Variable, error) {
	db := r.db.Order("id DESC")
	var items []Variable

	if len(params) > 0 {
		db = db.Where(params)
	}
	result := db.Find(&items)
	err := result.Error
	if err != nil {
		return items, errutil.NewGormError(err)
	}
	return items, err
}

func (r Repo) Retrieve(id int) (*Variable, error) {
	var item Variable
	result := r.db.Where("id = ?", id).First(&item)
	err := result.Error
	if err != nil {
		return &item, errutil.NewGormError(err)
	}
	return &item, err
}

func (r Repo) Create(item *Variable) (*Variable, error) {
	result := r.db.Create(item)
	err := result.Error
	if err != nil {
		return item, errutil.NewGormError(err)
	}
	return item, err
}

func (r Repo) Update(id int, data ctype.Dict) (*Variable, error) {
	item, err := r.Retrieve(id)
	if err != nil {
		return nil, err
	}
	result := r.db.Model(&item).Updates(data)
	err = result.Error
	if err != nil {
		return nil, errutil.NewGormError(err)
	}
	return item, err
}

func (r Repo) Delete(id int) ([]int, error) {
	ids := []int{id}
	result := r.db.Where("id = ?", id).Delete(&Variable{})
	err := result.Error
	if err != nil {
		return ids, errutil.NewGormError(err)
	}
	return ids, err
}

func (r Repo) DeleteList(ids []int) ([]int, error) {
	result := r.db.Where("id IN (?)", ids).Delete(&Variable{})
	err := result.Error
	if err != nil {
		return ids, errutil.NewGormError(err)
	}
	return ids, err
}
