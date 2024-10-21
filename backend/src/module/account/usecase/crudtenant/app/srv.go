package app

import (
	"src/common/ctype"
	"src/common/intf"
	"src/module/account/schema"
	"src/util/restlistutil"
)

type Schema = schema.Tenant

type Data struct {
	AuthClientID uint   `json:"auth_client_id" validate:"required"`
	Uid          string `json:"uid" validate:"required"`
	Title        string `json:"title" validate:"required"`
	Avatar       string `json:"avatar"`
	AvatarStr    string `json:"avatar_str"`
}

func (data Data) ToSchema() *Schema {
	return &Schema{
		AuthClientID: data.AuthClientID,
		Uid:          data.Uid,
		Title:        data.Title,
		Avatar:       data.Avatar,
		AvatarStr:    data.AvatarStr,
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
