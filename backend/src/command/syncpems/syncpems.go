package main

import (
	"src/command/syncpems/infra"
	"src/route"
	"src/util/dbutil"

	"github.com/labstack/echo/v4"
)

func main() {
	dbutil.InitDb()
	db := dbutil.Db()
	repo := infra.New(db)

	e := echo.New()
	apiGroup := e.Group("/api/v1")
	_, pemMap := route.CollectRoutes(apiGroup)

	repo.WritePems(pemMap)
	repo.EnsureTenantsRoles()
	repo.EnsureRolesPems(pemMap)
}
