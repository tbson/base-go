package route

import (
	"src/common/ctype"
	sso "src/module/account/usecase/auth/infra"
	crudauthclient "src/module/account/usecase/crudauthclient/infra"
	crudtenant "src/module/account/usecase/crudtenant/infra"
	cruduser "src/module/account/usecase/cruduser/infra"
	crudvariable "src/module/config/usecase/crudvariable/infra"

	"github.com/labstack/echo/v4"
)

func CollectRoutes(e *echo.Group) (*echo.Group, ctype.RoleMap) {
	roleMap := ctype.RoleMap{}
	e, roleMap = crudvariable.RegisterUrls(e, roleMap)
	e, roleMap = crudauthclient.RegisterUrls(e, roleMap)
	e, roleMap = crudtenant.RegisterUrls(e, roleMap)
	e, roleMap = cruduser.RegisterUrls(e, roleMap)
	e, roleMap = sso.RegisterUrls(e, roleMap)
	return e, roleMap
}
