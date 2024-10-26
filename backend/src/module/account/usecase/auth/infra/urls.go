package infra

import (
	"src/util/routeutil"

	"src/common/ctype"

	"github.com/labstack/echo/v4"
)

type RoleMap map[string][]string

func RegisterUrls(e *echo.Group, roleMap ctype.RoleMap) (*echo.Group, ctype.RoleMap) {
	g := e.Group("/account/auth")
	rr := routeutil.RegisterRoute(g, roleMap)

	rr("GET", "/sso/login/check/:tenantUid", CheckAuthUrl, []string{}, "")
	rr("GET", "/sso/login/:tenantUid", GetAuthUrl, []string{}, "")
	rr("GET", "/sso/logout/:tenantUid", GetLogoutUrl, []string{}, "")
	rr("GET", "/sso/callback", Callback, []string{}, "")
	rr("GET", "/sso/refresh-token", RefreshToken, []string{}, "")
	return e, roleMap
}
