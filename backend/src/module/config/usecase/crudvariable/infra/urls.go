package infra

import (
	"src/util/routeutil"

	"src/common/constant"
	"src/common/ctype"

	"github.com/labstack/echo/v4"
)

type RoleMap map[string][]string

func RegisterUrls(e *echo.Group, roleMap ctype.RoleMap) (*echo.Group, ctype.RoleMap) {
	g := e.Group("/config/variable")
	rr := routeutil.RegisterRoute(g, roleMap)

	rr("GET", "/", List, []string{constant.UsrTypeAdmin, constant.UsrTypeStaff}, "Get variable list")
	rr("GET", "/:id", Retrieve, []string{constant.UsrTypeAdmin, constant.UsrTypeStaff}, "Get variable detail")
	rr("POST", "/", Create, []string{constant.UsrTypeAdmin}, "Create variable")
	rr("PUT", "/:id", Update, []string{constant.UsrTypeAdmin}, "Update variable")
	rr("DELETE", "/:id", Delete, []string{constant.UsrTypeAdmin}, "Delete variable")
	rr("DELETE", "/", DeleteList, []string{constant.UsrTypeAdmin}, "Delete list variable")
	return e, roleMap
}
