package app

import (
	"src/common/ctype"
	"src/module/config/schema"
	"src/util/restlistutil"
)

type Schema = schema.Variable

type Service struct {
	cruder CRUDer[Schema]
	pager  Pager[Schema]
}

func (s Service) New(cruder CRUDer[Schema], pager Pager[Schema]) Service {
	return Service{cruder, pager}
}

func (s Service) NewForPaging(pager Pager[Schema]) Service {
	return Service{nil, pager}
}

func (s Service) NewForCruding(cruder CRUDer[Schema]) Service {
	return Service{cruder, nil}
}

func (srv Service) List(
	options restlistutil.ListOptions,
	searchableFields []string,
) (restlistutil.ListRestfulResult[Schema], error) {
	return srv.pager.Paging(options, searchableFields)
}

func (srv Service) Retrieve(queryOptions ctype.QueryOptions) (*Schema, error) {
	return srv.cruder.Retrieve(queryOptions)
}

func (srv Service) Create(inputData ctype.Dict) (*Schema, error) {
	return srv.cruder.Create(inputData)
}

func (srv Service) Update(id int, inputData ctype.Dict) (*Schema, error) {
	return srv.cruder.Update(id, inputData)
}

func (srv Service) Delete(id int) ([]int, error) {
	return srv.cruder.Delete(id)
}

func (srv Service) DeleteList(ids []int) ([]int, error) {
	return srv.cruder.DeleteList(ids)
}
