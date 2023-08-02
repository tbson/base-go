package util

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func homePage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}

func GetRoutes(e *echo.Echo) *echo.Echo {
	e.GET("/", homePage)
	return e
}
