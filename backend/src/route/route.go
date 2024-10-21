package route

import (
	"src/common/ctype"
	crudauthclient "src/module/account/usecase/crudauthclient/infra"
	crudtenant "src/module/account/usecase/crudtenant/infra"
	sso "src/module/account/usecase/sso/infra"
	crudvariable "src/module/config/usecase/crudvariable/infra"

	"github.com/labstack/echo/v4"
)

func CollectRoutes(e *echo.Group) (*echo.Group, ctype.RoleMap) {
	roleMap := ctype.RoleMap{}
	e, roleMap = crudvariable.RegisterUrls(e, roleMap)
	e, roleMap = crudauthclient.RegisterUrls(e, roleMap)
	e, roleMap = crudtenant.RegisterUrls(e, roleMap)
	e, roleMap = sso.RegisterUrls(e, roleMap)
	return e, roleMap
}
