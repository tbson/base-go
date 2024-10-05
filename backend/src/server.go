package main

import (
	"embed"
	"src/route"
	"src/util/dbutil"
	"src/util/localeutil"
	"src/util/vldtutil"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

//go:embed util/localeutil/locales/*
var localeFS embed.FS

func blankMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func languageMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accept := c.Request().Header.Get("Accept-Language")
		localizer := localeutil.Init(accept)
		c.Set("localizer", localizer)
		c.Set("lang", accept)
		return next(c)
	}
}

func main() {
	dbutil.InitDb()
	e := echo.New()
	e.Validator = &vldtutil.CustomValidator{Validator: validator.New()}
	e.Use(blankMiddleware)
	e.Use(languageMiddleware)
	apiGroup := e.Group("/api/v1")
	route.CollectRoutes(apiGroup)
	e.Logger.Fatal(e.Start("0.0.0.0:4000"))
}
