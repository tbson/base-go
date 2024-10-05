package infra

import (
	"net/http"
	"src/util/vldtutil"

	"github.com/labstack/echo/v4"
	// "src/module/config/repo"
	// "src/module/config/usecase/crudvariable/app"
)

type VariableData struct {
	Key         string `json:"key" validate:"required"`
	Value       string `json:"value"`
	Description string `json:"description"`
	DataType    string `json:"data_type" validate:"required,oneof=STRING INTEGER FLOAT BOOLEAN DATE DATETIME"`
}

// Cache the required fields at startup

func List(c echo.Context) error {
	return c.String(http.StatusOK, "List variable")
}

func Retrieve(c echo.Context) error {
	return c.String(http.StatusOK, "Retrieve variable")
}

func Create(c echo.Context) error {
	result, error := vldtutil.ValidatePayload(c, VariableData{})

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
