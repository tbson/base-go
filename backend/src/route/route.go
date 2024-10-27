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

func CollectRoutes(e *echo.Group) (*echo.Group, ctype.PemMap) {
	pemMap := ctype.PemMap{}
	e, pemMap = crudvariable.RegisterUrls(e, pemMap)
	e, pemMap = crudauthclient.RegisterUrls(e, pemMap)
	e, pemMap = crudtenant.RegisterUrls(e, pemMap)
	e, pemMap = cruduser.RegisterUrls(e, pemMap)
	e, pemMap = sso.RegisterUrls(e, pemMap)
	return e, pemMap
}
