package app

import (
	"src/common/ctype"
	"src/module/account/schema"
	"src/module/account/usecase/crudauthclient/app/intf"
	"src/util/restlistutil"
)

type Schema = schema.AuthClient

type Data struct {
	Uid         string `json:"uid" validate:"required"`
	Description string `json:"description"`
	Secret      string `json:"secret" validate:"required"`
	Partition   string `json:"partition" validate:"required"`
	Default     bool   `json:"default"`
}

func (data Data) ToSchema() *Schema {
	return &Schema{
		Uid:         data.Uid,
		Description: data.Description,
		Secret:      data.Secret,
		Partition:   data.Partition,
		Default:     data.Default,
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

func (srv Service) Retrieve(params ctype.Dict) (*Schema, error) {
	return srv.repo.Retrieve(params)
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
