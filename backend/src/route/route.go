package route

import (
	"src/common/ctype"
	auth "src/module/account/usecase/auth/infra"
	crudauthclient "src/module/account/usecase/crudauthclient/infra"
	crudrole "src/module/account/usecase/crudrole/infra"
	crudtenant "src/module/account/usecase/crudtenant/infra"
	cruduser "src/module/account/usecase/cruduser/infra"
	profile "src/module/account/usecase/profile/infra"
	crudvariable "src/module/config/usecase/crudvariable/infra"
	configOption "src/module/config/usecase/option/infra"

	"github.com/labstack/echo/v4"
)

func CollectRoutes(e *echo.Group) (*echo.Group, ctype.PemMap) {
	pemMap := ctype.PemMap{}
	e, pemMap = configOption.RegisterUrls(e, pemMap)
	e, pemMap = crudvariable.RegisterUrls(e, pemMap)
	e, pemMap = crudauthclient.RegisterUrls(e, pemMap)
	e, pemMap = crudtenant.RegisterUrls(e, pemMap)
	e, pemMap = crudrole.RegisterUrls(e, pemMap)
	e, pemMap = cruduser.RegisterUrls(e, pemMap)
	e, pemMap = auth.RegisterUrls(e, pemMap)
	e, pemMap = profile.RegisterUrls(e, pemMap)
	return e, pemMap
}
