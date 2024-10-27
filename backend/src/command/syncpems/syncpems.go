package main

import (
	"src/common/ctype"
	"src/module/account/repo/pem"
	"src/route"
	"src/util/dbutil"

	"github.com/labstack/echo/v4"
)

func main() {
	dbutil.InitDb()
	db := dbutil.Db()
	pemRepo := pem.New(db)

	e := echo.New()
	apiGroup := e.Group("/api/v1")
	_, roleMap := route.CollectRoutes(apiGroup)

	for _, value := range roleMap {
		data := ctype.Dict{
			"Title":  value.Title,
			"Module": value.Module,
			"Action": value.Action,
		}

		filterOptions := ctype.QueryOptions{
			Filters: ctype.Dict{
				"module": value.Module,
				"action": value.Action,
			},
		}

		_, err := pemRepo.GetOrCreate(filterOptions, data)

		if err != nil {
			panic(err)
		}
	}
}
