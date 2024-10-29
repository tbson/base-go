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

	rr.Public("GET", "/sso/login/check/:tenantUid", CheckAuthUrl)
	rr.Public("GET", "/sso/login/:tenantUid", GetAuthUrl)
	rr.Public("GET", "/sso/logout/:tenantUid", GetLogoutUrl)
	rr.Public("GET", "/sso/callback", Callback)
	rr.Public("GET", "/sso/post-logout", PostLogout)
	rr.Public("GET", "/sso/refresh-token", RefreshToken)
	rr.Private("GET", "/sso/refresh-token-check", RefreshTokenCheck)
	return e, pemMap
}
