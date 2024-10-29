package app

import (
	"src/common/ctype"
	"src/module/config/schema"
	"src/module/config/usecase/crudvariable/app/intf"
	"src/util/restlistutil"
)

type Schema = schema.Variable

type Service struct {
	cruder intf.CRUDer[Schema]
	pager  intf.Pager[Schema]
}

func (s Service) New(cruder intf.CRUDer[Schema], pager intf.Pager[Schema]) Service {
	return Service{cruder, pager}
}

func (s Service) NewForPaging(pager intf.Pager[Schema]) Service {
	return Service{nil, pager}
}

func (s Service) NewForCruding(cruder intf.CRUDer[Schema]) Service {
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
