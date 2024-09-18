package crudvariable

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	return c.String(http.StatusOK, "List variable")
}

func Retrieve(c echo.Context) error {
	return c.String(http.StatusOK, "Retrieve variable")
}

func Create(c echo.Context) error {
	return c.String(http.StatusOK, "Create variable")
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
