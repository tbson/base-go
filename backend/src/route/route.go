package route

import (
	"app/common/ctype"
	"app/module/config/usecase/crudvariable"

	"github.com/labstack/echo/v4"
)

func CollectRoutes(e *echo.Group) (*echo.Group, ctype.RoleMap) {
	roleMap := ctype.RoleMap{}
	e, roleMap = crudvariable.RegisterCtrl(e, roleMap)
	return e, roleMap
}
