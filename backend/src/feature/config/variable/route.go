package variable

import (
	ctrl "app/feature/config/variable/crud/ctrl"
	"reflect"
	"runtime"

	"app/util/constant"
	"app/util/ctype"

	"github.com/labstack/echo/v4"
)

type RoleMap map[string][]string

func ApplyRoutes(e *echo.Echo, roleMap ctype.RoleMap) (*echo.Echo, ctype.RoleMap) {
	g := e.Group("/config/variable")
	g.GET("/", ctrl.List)
	key := runtime.FuncForPC(reflect.ValueOf(ctrl.List).Pointer()).Name()
	roleMap[key] = []string{constant.UsrTypeAdmin, constant.UsrTypeStaff}

	g.POST("/", ctrl.Create)
	key1 := runtime.FuncForPC(reflect.ValueOf(ctrl.Create).Pointer()).Name()
	roleMap[key1] = []string{constant.UsrTypeAdmin}
	return e, roleMap
}
