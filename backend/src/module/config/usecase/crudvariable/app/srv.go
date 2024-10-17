package app

import (
	"src/common/ctype"
	"src/module/config/schema"
	"src/module/config/usecase/crudvariable/app/intf"
	"src/util/restlistutil"
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
	repo intf.RestCrudRepo[schema.Variable]
}

func NewCrudVariableSrv(repo intf.RestCrudRepo[schema.Variable]) CrudVariableSrv {
	return CrudVariableSrv{repo}
}

func (srv CrudVariableSrv) ListRestful(options restlistutil.ListOptions, searchableFields []string) (restlistutil.ListRestfulResult[schema.Variable], error) {
	return srv.repo.List(options, searchableFields)
}

func (srv CrudVariableSrv) RetrieveVariable(id int) (*schema.Variable, error) {
	return srv.repo.Retrieve(id)
}

func (srv CrudVariableSrv) CreateVariable(inputData VariableData) (*schema.Variable, error) {
	schema := inputData.ToSchema()
	return srv.repo.Create(schema)
}

func (srv CrudVariableSrv) UpdateVariable(id int, inputData ctype.Dict) (*schema.Variable, error) {
	return srv.repo.Update(id, inputData)
}

func (srv CrudVariableSrv) DeleteVariable(id int) ([]int, error) {
	return srv.repo.Delete(id)
}

func (srv CrudVariableSrv) DeleteListVariable(ids []int) ([]int, error) {
	return srv.repo.DeleteList(ids)
}
