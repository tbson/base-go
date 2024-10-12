package infra

import (
	"net/http"
	"src/util/vldtutil"

	"src/module/config/repo"
	"src/module/config/usecase/crudvariable/app"

	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	srv := app.NewCrudVariableSrv(repo.VariableRepo{})

	result, error := srv.ListVariable()

	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.JSON(http.StatusOK, result)
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

	return c.JSON(http.StatusOK, result)

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

	error := srv.DeleteVariable(id)

	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.String(http.StatusOK, "Delete variable")
}

func DeleteList(c echo.Context) error {
	ids := vldtutil.ValidateIds(c.QueryParam("ids"))

	srv := app.NewCrudVariableSrv(repo.VariableRepo{})
	error := srv.DeleteListVariable(ids)

	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.String(http.StatusOK, "Delete variables")
}
