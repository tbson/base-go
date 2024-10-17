package intf

import (
	"src/common/ctype"
	"src/module/config/schema"
	"src/util/restlistutil"
)

type VariableRepo interface {
	ListVariable(params ctype.Dict) ([]schema.Variable, error)
	RetrieveVariable(key int) (*schema.Variable, error)
	CreateVariable(variable *schema.Variable) (*schema.Variable, error)
	UpdateVariable(key int, variable ctype.Dict) (*schema.Variable, error)
	DeleteVariable(key int) ([]int, error)
	DeleteListVariable(keys []int) ([]int, error)
}

type VariableListRepo interface {
	ListRestful(
		options restlistutil.ListOptions,
		searchableFields []string) (restlistutil.ListRestfulResult[schema.Variable], error)
}
