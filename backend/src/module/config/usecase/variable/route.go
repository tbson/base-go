package variable

import (
	ctrl "app/module/config/usecase/variable/crud/ctrl"
	"app/util/routeutil"

	"app/common/constant"
	"app/common/ctype"

	"github.com/labstack/echo/v4"
)

type RoleMap map[string][]string

func RegisterCtrl(e *echo.Group, roleMap ctype.RoleMap) (*echo.Group, ctype.RoleMap) {
	g := e.Group("/config/variable")
	rr := routeutil.RegisterRoute(g, roleMap)

	rr("GET", "/", ctrl.List, []string{constant.UsrTypeAdmin, constant.UsrTypeStaff}, "Get variable list")
	rr("POST", "/", ctrl.Create, []string{constant.UsrTypeAdmin}, "Create variable list")

	return e, roleMap
}
