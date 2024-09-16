package main

import (
	"app/util/frameworkutil"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	apiGroup := e.Group("/api/v1")
	_, roleMap := frameworkutil.CollectRoutes(apiGroup)
	fmt.Println(roleMap)
}
