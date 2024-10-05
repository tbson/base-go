package core

import (
	"src/module/config/repo/schema"
)

type VariablePayload struct {
	Key         string `json:"key" validate:"required"`
	Value       string `json:"value"`
	Description string `json:"description"`
	DataType    string `json:"data_type" validate:"required"`
}

type CrudVariableSrv struct {
	repo VariableRepo
}

func (s *CrudVariableSrv) CreateVariable(variable *schema.Variable) (*schema.Variable, error) {
	return s.repo.CreateVariable(variable)
}

func NewCrudVariableSrv(repo VariableRepo) *CrudVariableSrv {
	return &CrudVariableSrv{repo}
}
