package infra

import (
	"src/util/routeutil"

	"src/common/ctype"
	"src/common/profiletype"

	"github.com/labstack/echo/v4"
)

type RoleMap map[string][]string

func RegisterUrls(e *echo.Group, pemMap ctype.PemMap) (*echo.Group, ctype.PemMap) {
	g := e.Group("/account/tenant")
	rr := routeutil.RegisterRoute(g, pemMap)

	rr.Rbac("GET", "/", List, []string{profiletype.ADMIN, profiletype.STAFF}, "Get tenant list")
	rr.Rbac("GET", "/:id", Retrieve, []string{profiletype.ADMIN, profiletype.STAFF}, "Get tenant detail")
	rr.Rbac("POST", "/", Create, []string{profiletype.ADMIN, profiletype.STAFF}, "Create tenant")
	rr.Rbac("PUT", "/:id", Update, []string{profiletype.ADMIN, profiletype.STAFF, profiletype.MANAGER}, "Update tenant")
	rr.Rbac("DELETE", "/:id", Delete, []string{profiletype.ADMIN, profiletype.STAFF}, "Delete tenant")
	rr.Rbac("DELETE", "/", DeleteList, []string{profiletype.ADMIN, profiletype.STAFF}, "Delete list tenant")
	return e, pemMap
}
