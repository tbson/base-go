package infra

import (
	"src/util/routeutil"

	"src/common/ctype"

	"github.com/labstack/echo/v4"
)

type RoleMap map[string][]string

func RegisterUrls(e *echo.Group, roleMap ctype.RoleMap) (*echo.Group, ctype.RoleMap) {
	g := e.Group("/account/sso")
	rr := routeutil.RegisterRoute(g, roleMap)

	rr("GET", "/login-url/:tenantId", GetLoginUrl, []string{}, "")
	rr("GET", "/callback", Callback, []string{}, "")
	return e, roleMap
}
