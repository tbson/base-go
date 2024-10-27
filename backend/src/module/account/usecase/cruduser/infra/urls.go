package infra

import (
	"src/util/routeutil"

	"src/common/ctype"
	"src/common/profiletype"

	"github.com/labstack/echo/v4"
)

type RoleMap map[string][]string

func RegisterUrls(e *echo.Group, pemMap ctype.PemMap) (*echo.Group, ctype.PemMap) {
	g := e.Group("/account/user")
	rr := routeutil.RegisterRoute(g, pemMap)

	rr("GET", "/", List, []string{profiletype.ADMIN, profiletype.STAFF, profiletype.MANAGER}, "Get user list")
	rr("GET", "/:id", Retrieve, []string{profiletype.ADMIN, profiletype.STAFF, profiletype.MANAGER}, "Get user detail")
	rr("POST", "/", Create, []string{profiletype.ADMIN, profiletype.STAFF, profiletype.MANAGER}, "Create user")
	rr("PUT", "/:id", Update, []string{profiletype.ADMIN, profiletype.STAFF, profiletype.MANAGER}, "Update user")
	rr("DELETE", "/:id", Delete, []string{profiletype.ADMIN, profiletype.STAFF, profiletype.MANAGER}, "Delete user")
	rr("DELETE", "/", DeleteList, []string{profiletype.ADMIN, profiletype.STAFF, profiletype.MANAGER}, "Delete list user")
	return e, pemMap
}
