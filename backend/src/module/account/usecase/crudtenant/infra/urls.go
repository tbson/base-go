package infra

import (
	"src/util/routeutil"

	"src/common/ctype"
	"src/common/profiletype"

	"github.com/labstack/echo/v4"
)

type RoleMap map[string][]string

func RegisterUrls(e *echo.Group, roleMap ctype.RoleMap) (*echo.Group, ctype.RoleMap) {
	g := e.Group("/account/tenant")
	rr := routeutil.RegisterRoute(g, roleMap)

	rr("GET", "/", List, []string{profiletype.ADMIN, profiletype.STAFF}, "Get tenant list")
	rr("GET", "/:id", Retrieve, []string{profiletype.ADMIN, profiletype.STAFF}, "Get tenant detail")
	rr("POST", "/", Create, []string{profiletype.ADMIN, profiletype.STAFF}, "Create tenant")
	rr("PUT", "/:id", Update, []string{profiletype.ADMIN, profiletype.STAFF, profiletype.MANAGER}, "Update tenant")
	rr("DELETE", "/:id", Delete, []string{profiletype.ADMIN, profiletype.STAFF}, "Delete tenant")
	rr("DELETE", "/", DeleteList, []string{profiletype.ADMIN, profiletype.STAFF}, "Delete list tenant")
	return e, roleMap
}
