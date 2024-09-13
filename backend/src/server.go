package main

import (
	"app/route"
	"app/util/dbutil"

	"github.com/labstack/echo/v4"
)

func blankMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func main() {
	dbutil.InitDb()
	e := echo.New()
	apiGroup := e.Group("/api/v1")
	route.CollectRoutes(apiGroup)
	e.Logger.Fatal(e.Start("0.0.0.0:4000"))
}
