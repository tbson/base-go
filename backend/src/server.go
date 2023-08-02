package main

import (
	"app/feature/config/variable"
	"fmt"
	"net/http"

	"github.com/samber/lo"

	"app/util/ctype"

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
	e := echo.New()
	roleMap := ctype.RoleMap{}
	e, roleMap = variable.ApplyRoutes(e, roleMap)
	fmt.Println(roleMap)
	verbs := []string{"GET", "POST", "PUT", "DELETE"}
	for _, route := range e.Routes() {
		// check if route.Method in verbs
		if lo.Contains(verbs, route.Method) {
			fmt.Println(route.Method)
			fmt.Println(route.Path)
			fmt.Println(route.Name)
			fmt.Println("---------------")
		}
	}
	e.Logger.Fatal(e.Start("0.0.0.0:4000"))
}

func GetRoles() []string {
	return []string{"Admin", "User"}
}
