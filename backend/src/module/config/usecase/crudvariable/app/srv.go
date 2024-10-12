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

func (data VariableData) ToSchema() *schema.Variable {
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

func NewCrudVariableSrv(repo intf.VariableRepo) CrudVariableSrv {
	return CrudVariableSrv{repo}
}

func (srv CrudVariableSrv) ListVariable() ([]schema.Variable, error) {
	return srv.repo.ListVariable()
}

func (srv CrudVariableSrv) RetrieveVariable(id int) (*schema.Variable, error) {
	return srv.repo.RetrieveVariable(id)
}

func (srv CrudVariableSrv) CreateVariable(inputData VariableData) (*schema.Variable, error) {
	schema := inputData.ToSchema()
	return srv.repo.CreateVariable(schema)
}

func (srv CrudVariableSrv) UpdateVariable(id int, inputData map[string]interface{}) (*schema.Variable, error) {
	return srv.repo.UpdateVariable(id, inputData)
}

func (srv CrudVariableSrv) DeleteVariable(id int) error {
	return srv.repo.DeleteVariable(id)
}

func (srv CrudVariableSrv) DeleteListVariable(ids []int) error {
	return srv.repo.DeleteListVariable(ids)
}
