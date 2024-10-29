package app

import (
	"src/common/ctype"
	"src/util/restlistutil"
)

type Pager[T any] interface {
	Paging(
		options restlistutil.ListOptions,
		searchableFields []string) (
		restlistutil.ListRestfulResult[T], error,
	)
}

type CRUDer[T any] interface {
	Retrieve(queryOptions ctype.QueryOptions) (*T, error)
	Create(data ctype.Dict) (*T, error)
	Update(key int, data ctype.Dict) (*T, error)
	Delete(key int) ([]int, error)
	DeleteList(keys []int) ([]int, error)
}
