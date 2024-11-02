package main

import (
	"html/template"
	"io"
	"src/middleware"
	"src/route"
	"src/util/dbutil"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// CustomValidator implements the echo.Validator interface
type customValidator struct {
	Validator *validator.Validate
}

// Validate validates the input struct using the validator
func (cv *customValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func main() {
	dbutil.InitDb()
	e := echo.New()
	e.Debug = true
	e.Validator = &customValidator{Validator: validator.New()}
	t := &Template{
		templates: template.Must(template.ParseGlob("/code/public/views/*.html")),
	}
	e.Renderer = t

	e.Use(middleware.LangMiddleware)
	apiGroup := e.Group("/api/v1")
	route.CollectRoutes(apiGroup)
	e.Logger.Fatal(e.Start("0.0.0.0:4000"))
}
