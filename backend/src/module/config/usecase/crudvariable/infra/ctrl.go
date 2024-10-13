package infra

import (
	"net/http"
	"src/util/dbutil"
	"src/util/vldtutil"

	"src/module/config/repo"
	"src/module/config/usecase/crudvariable/app"

	"github.com/labstack/echo/v4"
)

var searchableFields = []string{"key", "value", "description"}
var filterableFields = []string{"data_type"}
var orderableFields = []string{"id", "key"}

func List(c echo.Context) error {
	srv := app.NewCrudVariableListSrv(VariableListRepo{})

	options := dbutil.GetOptions(c, filterableFields, orderableFields)
	listResult, error := srv.ListRestful(options, searchableFields)
	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.JSON(http.StatusOK, listResult)
}

func Retrieve(c echo.Context) error {
	srv := app.NewCrudVariableSrv(repo.VariableRepo{})
	id := vldtutil.ValidateId(c.Param("id"))
	result, error := srv.RetrieveVariable(id)

	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.JSON(http.StatusOK, result)
}

func Create(c echo.Context) error {

	data, error := vldtutil.ValidatePayload(c, app.VariableData{})

	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	srv := app.NewCrudVariableSrv(repo.VariableRepo{})

	result, error := srv.CreateVariable(data)

	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.JSON(http.StatusCreated, result)

}

func Update(c echo.Context) error {
	data, error := vldtutil.ValidateUpdatePayload(c)

	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	srv := app.NewCrudVariableSrv(repo.VariableRepo{})
	id := vldtutil.ValidateId(c.Param("id"))

	result, error := srv.UpdateVariable(id, data)

	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.JSON(http.StatusOK, result)
}

func Delete(c echo.Context) error {
	srv := app.NewCrudVariableSrv(repo.VariableRepo{})
	id := vldtutil.ValidateId(c.Param("id"))

	ids, error := srv.DeleteVariable(id)

	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}
	// return deleted id slide with one item
	return c.JSON(http.StatusNoContent, ids)
}

func DeleteList(c echo.Context) error {
	ids := vldtutil.ValidateIds(c.QueryParam("ids"))

	srv := app.NewCrudVariableSrv(repo.VariableRepo{})
	ids, error := srv.DeleteListVariable(ids)

	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}
	// return deleted id slide
	return c.JSON(http.StatusNoContent, ids)
}
