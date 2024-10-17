package app

import (
	"src/common/ctype"
	"src/module/config/schema"
	"src/module/config/usecase/crudvariable/app/intf"
	"src/util/restlistutil"
)

type Schema = schema.Variable

type Data struct {
	Key         string `json:"key" validate:"required"`
	Value       string `json:"value"`
	Description string `json:"description"`
	DataType    string `json:"data_type" validate:"required,oneof=STRING INTEGER FLOAT BOOLEAN DATE DATETIME"`
}

func (data Data) ToSchema() *Schema {
	return &Schema{
		Key:         data.Key,
		Value:       data.Value,
		Description: data.Description,
		DataType:    data.DataType,
	}
}

type Service struct {
	repo intf.RestCrudRepo[Schema]
}

func (s Service) New(repo intf.RestCrudRepo[Schema]) Service {
	return Service{repo}
}

func (srv Service) List(options restlistutil.ListOptions, searchableFields []string) (restlistutil.ListRestfulResult[Schema], error) {
	return srv.repo.List(options, searchableFields)
}

func (srv Service) Retrieve(id int) (*Schema, error) {
	return srv.repo.Retrieve(id)
}

func (srv Service) Create(inputData Data) (*Schema, error) {
	schema := inputData.ToSchema()
	return srv.repo.Create(schema)
}

func (srv Service) Update(id int, inputData ctype.Dict) (*Schema, error) {
	return srv.repo.Update(id, inputData)
}

func (srv Service) Delete(id int) ([]int, error) {
	return srv.repo.Delete(id)
}

func (srv Service) DeleteList(ids []int) ([]int, error) {
	return srv.repo.DeleteList(ids)
}
