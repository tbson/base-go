package app

import (
	"src/module/config/schema"
	"src/module/config/usecase/crudvariable/app/intf"
)

type VariableData struct {
	Key         string `json:"key" validate:"required"`
	Value       string `json:"value"`
	Description string `json:"description"`
	DataType    string `json:"data_type" validate:"required,oneof=STRING INTEGER FLOAT BOOLEAN DATE DATETIME"`
}

func NewVariable(data VariableData) *schema.Variable {
	return &schema.Variable{
		Key:         data.Key,
		Value:       data.Value,
		Description: data.Description,
		DataType:    data.DataType,
	}
}

type CrudVariableSrv struct {
	repo intf.VariableRepo
}

func (srv CrudVariableSrv) CreateVariable(inputData VariableData) (*schema.Variable, error) {
	variable := NewVariable(inputData)
	return srv.repo.CreateVariable(variable)
}

func NewCrudVariableSrv(repo intf.VariableRepo) CrudVariableSrv {
	return CrudVariableSrv{repo}
}
