package infra

import (
	"src/util/routeutil"

	"src/common/constant"
	"src/common/ctype"

	"github.com/labstack/echo/v4"
)

type RoleMap map[string][]string

func RegisterUrls(e *echo.Group, roleMap ctype.RoleMap) (*echo.Group, ctype.RoleMap) {
	g := e.Group("/account/auth-client")
	rr := routeutil.RegisterRoute(g, roleMap)

	rr("GET", "/", List, []string{constant.UsrTypeAdmin, constant.UsrTypeStaff}, "Get auth client list")
	rr("GET", "/:id", Retrieve, []string{constant.UsrTypeAdmin, constant.UsrTypeStaff}, "Get auth client detail")
	rr("POST", "/", Create, []string{constant.UsrTypeAdmin}, "Create auth client")
	rr("PUT", "/:id", Update, []string{constant.UsrTypeAdmin}, "Update auth client")
	rr("DELETE", "/:id", Delete, []string{constant.UsrTypeAdmin}, "Delete auth client")
	rr("DELETE", "/", DeleteList, []string{constant.UsrTypeAdmin}, "Delete list auth client")
	return e, roleMap
}
