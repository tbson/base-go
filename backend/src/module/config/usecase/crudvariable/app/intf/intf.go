package intf

import (
	"src/common/ctype"
	"src/util/restlistutil"
)

type RestCrudRepo[T any] interface {
	List(
		options restlistutil.ListOptions,
		searchableFields []string) (
		restlistutil.ListRestfulResult[T], error,
	)
	Retrieve(key int) (*T, error)
	Create(instance *T) (*T, error)
	Update(key int, data ctype.Dict) (*T, error)
	Delete(key int) ([]int, error)
	DeleteList(keys []int) ([]int, error)
}
