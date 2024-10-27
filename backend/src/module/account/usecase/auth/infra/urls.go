package infra

import (
	"src/util/routeutil"

	"src/common/ctype"

	"github.com/labstack/echo/v4"
)

type RoleMap map[string][]string

func RegisterUrls(e *echo.Group, pemMap ctype.PemMap) (*echo.Group, ctype.PemMap) {
	g := e.Group("/account/auth")
	rr := routeutil.RegisterRoute(g, pemMap)

	rr("GET", "/sso/login/check/:tenantUid", CheckAuthUrl, []string{}, "")
	rr("GET", "/sso/login/:tenantUid", GetAuthUrl, []string{}, "")
	rr("GET", "/sso/logout/:tenantUid", GetLogoutUrl, []string{}, "")
	rr("GET", "/sso/callback", Callback, []string{}, "")
	rr("GET", "/sso/refresh-token", RefreshToken, []string{}, "")
	return e, pemMap
}
