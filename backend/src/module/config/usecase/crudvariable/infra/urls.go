package infra

import (
	"src/util/routeutil"

	"src/common/ctype"
	"src/common/profiletype"

	"github.com/labstack/echo/v4"
)

type RoleMap map[string][]string

func RegisterUrls(e *echo.Group, pemMap ctype.PemMap) (*echo.Group, ctype.PemMap) {
	g := e.Group("/config/variable")
	rr := routeutil.RegisterRoute(g, pemMap)

	rr.Rbac("GET", "/", List, []string{profiletype.ADMIN, profiletype.STAFF}, "Get variable list")
	rr.Rbac("GET", "/:id", Retrieve, []string{profiletype.ADMIN, profiletype.STAFF}, "Get variable detail")
	rr.Rbac("POST", "/", Create, []string{profiletype.ADMIN}, "Create variable")
	rr.Rbac("PUT", "/:id", Update, []string{profiletype.ADMIN}, "Update variable")
	rr.Rbac("DELETE", "/:id", Delete, []string{profiletype.ADMIN}, "Delete variable")
	rr.Rbac("DELETE", "/", DeleteList, []string{profiletype.ADMIN}, "Delete list variable")
	return e, pemMap
}
