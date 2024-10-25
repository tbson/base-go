package cookieutil

import (
	"net/http"
	"src/common/setting"
)

func newCookie(name string, value string, path string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Domain = setting.DOMAIN         // Set the Domain attribute
	cookie.Path = path                     // Set the Path attribute
	cookie.Secure = true                   // Set the Secure attribute
	cookie.HttpOnly = true                 // Prevents JavaScript access (optional)
	cookie.SameSite = http.SameSiteLaxMode // Set the SameSite attribute
	return cookie
}

func NewAccessTokenCookie(value string) *http.Cookie {
	return newCookie("access_token", value, "/api/v1/")
}

func NewRefreshTokenCookie(value string) *http.Cookie {
	return newCookie("refresh_token", value, "/api/v1/account/auth/sso/refresh-token")
}
