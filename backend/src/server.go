package main

import (
	"html/template"
	"io"
	"src/route"
	"src/util/dbutil"
	"src/util/localeutil"
	"src/util/vldtutil"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

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
	e.Debug = true
	e.Validator = &vldtutil.CustomValidator{Validator: validator.New()}
	t := &Template{
		templates: template.Must(template.ParseGlob("/code/public/views/*.html")),
	}
	e.Renderer = t

	e.Use(blankMiddleware)
	e.Use(languageMiddleware)
	apiGroup := e.Group("/api/v1")
	route.CollectRoutes(apiGroup)
	e.Logger.Fatal(e.Start("0.0.0.0:4000"))
}
