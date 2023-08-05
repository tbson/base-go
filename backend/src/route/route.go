package route

import (
	"app/common/ctype"
	"app/feature/config/variable"

	"github.com/labstack/echo/v4"
)

func ApplyRoutes(e *echo.Echo) (*echo.Echo, ctype.RoleMap) {
	roleMap := ctype.RoleMap{}
	e, roleMap = variable.ApplyRoutes(e, roleMap)
	return e, roleMap
}
