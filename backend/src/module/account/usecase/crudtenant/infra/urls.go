package infra

import (
	"src/util/routeutil"

	"src/common/constant"
	"src/common/ctype"

	"github.com/labstack/echo/v4"
)

type RoleMap map[string][]string

func RegisterUrls(e *echo.Group, roleMap ctype.RoleMap) (*echo.Group, ctype.RoleMap) {
	g := e.Group("/account/tenant")
	rr := routeutil.RegisterRoute(g, roleMap)

	rr("GET", "/", List, []string{constant.ProfileTypeAdmin}, "Get auth client list")
	rr("GET", "/:id", Retrieve, []string{constant.ProfileTypeAdmin}, "Get auth client detail")
	rr("POST", "/", Create, []string{constant.ProfileTypeAdmin}, "Create auth client")
	rr("PUT", "/:id", Update, []string{constant.ProfileTypeAdmin, constant.ProfileTypeManager}, "Update auth client")
	rr("DELETE", "/:id", Delete, []string{constant.ProfileTypeAdmin}, "Delete auth client")
	rr("DELETE", "/", DeleteList, []string{constant.ProfileTypeAdmin}, "Delete list auth client")
	return e, roleMap
}
