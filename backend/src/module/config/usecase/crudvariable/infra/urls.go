package infra

import (
	"src/util/routeutil"

	"src/common/ctype"
	"src/common/profiletype"

	"github.com/labstack/echo/v4"
)

type RoleMap map[string][]string

func RegisterUrls(e *echo.Group, roleMap ctype.RoleMap) (*echo.Group, ctype.RoleMap) {
	g := e.Group("/config/variable")
	rr := routeutil.RegisterRoute(g, roleMap)

	rr("GET", "/", List, []string{profiletype.ADMIN, profiletype.STAFF}, "Get variable list")
	rr("GET", "/:id", Retrieve, []string{profiletype.ADMIN, profiletype.STAFF}, "Get variable detail")
	rr("POST", "/", Create, []string{profiletype.ADMIN}, "Create variable")
	rr("PUT", "/:id", Update, []string{profiletype.ADMIN}, "Update variable")
	rr("DELETE", "/:id", Delete, []string{profiletype.ADMIN}, "Delete variable")
	rr("DELETE", "/", DeleteList, []string{profiletype.ADMIN}, "Delete list variable")
	return e, roleMap
}
