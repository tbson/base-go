package main

import (
	"net/http"

	"app/route"

	"app/util/db_util"

	"github.com/labstack/echo/v4"
)

func blankMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Middleware logic (can be empty)
		return next(c)
	}
}

func homePage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}

func otherPage(c echo.Context) error {
	return c.String(http.StatusOK, "Other page")
}

func main() {
	db_util.InitDb()
	e := echo.New()
	e, _ = route.ApplyRoutes(e)
	e.Logger.Fatal(e.Start("0.0.0.0:4000"))
}

func GetRoles() []string {
	return []string{"Admin", "User"}
}
