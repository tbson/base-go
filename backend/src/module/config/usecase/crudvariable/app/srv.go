package app

import (
	"src/util/restlistutil"
)

type Service[T any] struct {
	pager Pager[T]
}

func New[T any](pager Pager[T]) Service[T] {
	return Service[T]{pager}
}

func (srv Service[T]) List(
	options restlistutil.ListOptions,
	searchableFields []string,
) (restlistutil.ListRestfulResult[T], error) {
	return srv.pager.Paging(options, searchableFields)
}
