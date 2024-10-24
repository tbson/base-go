package user

import (
	"src/common/ctype"
	"src/module/account/schema"
	"src/util/errutil"

	"gorm.io/gorm"
)

type Schema = schema.User

type Repo struct {
	client *gorm.DB
}

func New(client *gorm.DB) Repo {
	return Repo{client: client}
}

func (r Repo) List(queryOptions ctype.QueryOptions) ([]Schema, error) {
	db := r.client
	db = db.Order("id DESC")
	filters := queryOptions.Filters
	preloads := queryOptions.Preloads
	if len(preloads) > 0 {
		for _, preload := range preloads {
			db = db.Preload(preload)
		}
	}

	var items []Schema

	if len(filters) > 0 {
		db = db.Where(map[string]interface{}(filters))
	}
	result := db.Find(&items)
	err := result.Error
	if err != nil {
		return items, errutil.NewGormError(err)
	}
	return items, err
}

func (r Repo) Retrieve(queryOptions ctype.QueryOptions) (*Schema, error) {
	db := r.client
	filters := queryOptions.Filters
	preloads := queryOptions.Preloads
	if len(preloads) > 0 {
		for _, preload := range preloads {
			db = db.Preload(preload)
		}
	}

	var item Schema
	result := r.client.Where(map[string]interface{}(filters)).First(&item)
	err := result.Error
	if err != nil {
		return &item, errutil.NewGormError(err)
	}
	return &item, err
}

func (r Repo) Create(item *Schema) (*Schema, error) {
	result := r.client.Create(item)
	err := result.Error
	if err != nil {
		return item, errutil.NewGormError(err)
	}
	return item, err
}

func (r Repo) GetOrCreate(queryOptions ctype.QueryOptions, item *Schema) (*Schema, error) {
	existItem, err := r.Retrieve(queryOptions)
	if err != nil {
		return r.Create(item)
	}
	return existItem, nil
}

func (r Repo) Update(id int, data ctype.Dict) (*Schema, error) {
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"id": id},
	}
	item, err := r.Retrieve(queryOptions)
	if err != nil {
		return nil, err
	}
	result := r.client.Model(&item).Updates(map[string]interface{}(data))
	err = result.Error
	if err != nil {
		return nil, errutil.NewGormError(err)
	}
	return item, err
}

func (r Repo) Delete(id int) ([]int, error) {
	ids := []int{id}
	result := r.client.Where("id = ?", id).Delete(&Schema{})
	err := result.Error
	if err != nil {
		return ids, errutil.NewGormError(err)
	}
	return ids, err
}

func (r Repo) DeleteList(ids []int) ([]int, error) {
	result := r.client.Where("id IN (?)", ids).Delete(&Schema{})
	err := result.Error
	if err != nil {
		return ids, errutil.NewGormError(err)
	}
	return ids, err
}
