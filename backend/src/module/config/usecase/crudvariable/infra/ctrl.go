package infra

import (
	"net/http"
	"src/util/vldtutil"

	"src/module/config/repo"
	"src/module/config/usecase/crudvariable/app"

	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	return c.String(http.StatusOK, "List variable")
}

func Retrieve(c echo.Context) error {
	return c.String(http.StatusOK, "Retrieve variable")
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
	return c.String(http.StatusOK, "Update variable")
}

func Delete(c echo.Context) error {
	return c.String(http.StatusOK, "Delete variable")
}

func DeleteList(c echo.Context) error {
	return c.String(http.StatusOK, "Delete list variable")
}
