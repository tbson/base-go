package route

import (
	"src/common/ctype"
	crudvariable "src/module/config/usecase/crudvariable/infra"

	"github.com/labstack/echo/v4"
)

func CollectRoutes(e *echo.Group) (*echo.Group, ctype.RoleMap) {
	roleMap := ctype.RoleMap{}
	e, roleMap = crudvariable.RegisterUrls(e, roleMap)
	return e, roleMap
}
