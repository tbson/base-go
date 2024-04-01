package main

import (
	"app/route"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	apiGroup := e.Group("/api/v1")
	_, roleMap := route.CollectRoutes(apiGroup)
	fmt.Println(roleMap)
}
