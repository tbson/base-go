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

func (s CrudVariableSrv) New(repo intf.RestCrudRepo[schema.Variable]) CrudVariableSrv {
	return CrudVariableSrv{repo}
}

func (srv CrudVariableSrv) List(options restlistutil.ListOptions, searchableFields []string) (restlistutil.ListRestfulResult[schema.Variable], error) {
	return srv.repo.List(options, searchableFields)
}

func (srv CrudVariableSrv) Retrieve(id int) (*schema.Variable, error) {
	return srv.repo.Retrieve(id)
}

func (srv CrudVariableSrv) Create(inputData VariableData) (*schema.Variable, error) {
	schema := inputData.ToSchema()
	return srv.repo.Create(schema)
}

func (srv CrudVariableSrv) Update(id int, inputData ctype.Dict) (*schema.Variable, error) {
	return srv.repo.Update(id, inputData)
}

func (srv CrudVariableSrv) Delete(id int) ([]int, error) {
	return srv.repo.Delete(id)
}

func (srv CrudVariableSrv) DeleteList(ids []int) ([]int, error) {
	return srv.repo.DeleteList(ids)
}
