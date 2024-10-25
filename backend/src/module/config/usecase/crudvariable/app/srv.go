package app

import (
	"src/common/ctype"
	"src/module/config/schema"
	"src/module/config/usecase/crudvariable/app/intf"
	"src/util/restlistutil"
)

type Schema = schema.Variable

type Service struct {
	repo intf.RestCrudRepo[Schema]
}

func (s Service) New(repo intf.RestCrudRepo[Schema]) Service {
	return Service{repo}
}

func (srv Service) List(
	options restlistutil.ListOptions,
	searchableFields []string,
) (restlistutil.ListRestfulResult[Schema], error) {
	return srv.repo.List(options, searchableFields)
}

func (srv Service) Retrieve(queryOptions ctype.QueryOptions) (*Schema, error) {
	return srv.repo.Retrieve(queryOptions)
}

func (srv Service) Create(inputData ctype.Dict) (*Schema, error) {
	return srv.repo.Create(inputData)
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
