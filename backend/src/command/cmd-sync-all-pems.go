package main

import (
	"app/route"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main1() {
	e := echo.New()
	_, roleMap := route.ApplyRoutes(e)
	fmt.Println(roleMap)
}
