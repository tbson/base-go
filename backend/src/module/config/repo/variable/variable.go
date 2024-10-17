package variable

import (
	"src/common/ctype"
	"src/module/config/schema"
	"src/util/errutil"

	"gorm.io/gorm"
)

type Schema = schema.Variable

type Repo struct {
	db *gorm.DB
}

func (r Repo) New(db *gorm.DB) Repo {
	return Repo{db: db}
}

func (r Repo) List(params ctype.Dict) ([]Schema, error) {
	db := r.db.Order("id DESC")
	var items []Schema

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

func (r Repo) Retrieve(id int) (*Schema, error) {
	var item Schema
	result := r.db.Where("id = ?", id).First(&item)
	err := result.Error
	if err != nil {
		return &item, errutil.NewGormError(err)
	}
	return &item, err
}

func (r Repo) Create(item *Schema) (*Schema, error) {
	result := r.db.Create(item)
	err := result.Error
	if err != nil {
		return item, errutil.NewGormError(err)
	}
	return item, err
}

func (r Repo) Update(id int, data ctype.Dict) (*Schema, error) {
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
	result := r.db.Where("id = ?", id).Delete(&Schema{})
	err := result.Error
	if err != nil {
		return ids, errutil.NewGormError(err)
	}
	return ids, err
}

func (r Repo) DeleteList(ids []int) ([]int, error) {
	result := r.db.Where("id IN (?)", ids).Delete(&Schema{})
	err := result.Error
	if err != nil {
		return ids, errutil.NewGormError(err)
	}
	return ids, err
}
