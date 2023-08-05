package variable

import (
	ctrl "app/feature/config/variable/crud/ctrl"
	"app/util/route_util"

	"app/common/constant"
	"app/common/ctype"

	"github.com/labstack/echo/v4"
)

type RoleMap map[string][]string

func ApplyRoutes(e *echo.Echo, roleMap ctype.RoleMap) (*echo.Echo, ctype.RoleMap) {
	g := e.Group("/config/variable")
	applyRoute := route_util.ApplyRoute(g, roleMap)

	applyRoute("GET", "/", ctrl.List, []string{constant.UsrTypeAdmin, constant.UsrTypeStaff}, "Get variable list")
	applyRoute("POST", "/", ctrl.Create, []string{constant.UsrTypeAdmin}, "Create variable list")

	return e, roleMap
}
