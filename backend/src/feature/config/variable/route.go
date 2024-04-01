package variable

import (
	ctrl "app/feature/config/variable/crud/ctrl"
	"app/util/route_util"

	"app/common/constant"
	"app/common/ctype"

	"github.com/labstack/echo/v4"
)

type RoleMap map[string][]string

func RegisterCtrl(e *echo.Group, roleMap ctype.RoleMap) (*echo.Group, ctype.RoleMap) {
	g := e.Group("/config/variable")
	rr := route_util.RegisterRoute(g, roleMap)

	rr("GET", "/", ctrl.List, []string{constant.UsrTypeAdmin, constant.UsrTypeStaff}, "Get variable list")
	rr("POST", "/", ctrl.Create, []string{constant.UsrTypeAdmin}, "Create variable list")

	return e, roleMap
}
