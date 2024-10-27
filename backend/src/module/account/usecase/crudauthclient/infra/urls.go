package infra

import (
	"src/util/routeutil"

	"src/common/ctype"
	"src/common/profiletype"

	"github.com/labstack/echo/v4"
)

type RoleMap map[string][]string

func RegisterUrls(e *echo.Group, pemMap ctype.PemMap) (*echo.Group, ctype.PemMap) {
	g := e.Group("/account/auth-client")
	rr := routeutil.RegisterRoute(g, pemMap)

	rr("GET", "/", List, []string{profiletype.ADMIN}, "Get auth client list")
	rr("GET", "/:id", Retrieve, []string{profiletype.ADMIN}, "Get auth client detail")
	rr("POST", "/", Create, []string{profiletype.ADMIN}, "Create auth client")
	rr("PUT", "/:id", Update, []string{profiletype.ADMIN}, "Update auth client")
	rr("DELETE", "/:id", Delete, []string{profiletype.ADMIN}, "Delete auth client")
	rr("DELETE", "/", DeleteList, []string{profiletype.ADMIN}, "Delete list auth client")
	return e, pemMap
}
