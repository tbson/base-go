package route

import (
	"app/common/ctype"
	"app/feature/config/variable"

	"github.com/labstack/echo/v4"
)

func CollectRoutes(e *echo.Group) (*echo.Group, ctype.RoleMap) {
	roleMap := ctype.RoleMap{}
	e, roleMap = variable.RegisterCtrl(e, roleMap)
	return e, roleMap
}
