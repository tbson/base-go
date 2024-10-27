package cookieutil

import (
	"net/http"
	"src/common/setting"
	"strings"

	"github.com/labstack/echo/v4"
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

func NewRealmCookie(value string) *http.Cookie {
	return newCookie("realm", value, "/api/v1/")
}

func NewRefreshTokenCookie(value string) *http.Cookie {
	return newCookie("refresh_token", value, "/api/v1/account/auth/sso/refresh-token")
}

func GetAccessToken(c echo.Context) (string, error) {
	cookieToken, err := c.Cookie("access_token")
	if err == nil {
		return cookieToken.Value, nil
	}

	headerToken := c.Request().Header.Get("Authorization")
	if headerToken != "" {
		return strings.Split(headerToken, " ")[1], nil
	}

	return "", err
}

func GetRealm(c echo.Context) (string, error) {
	cookieRealm, err := c.Cookie("realm")
	if err == nil {
		return cookieRealm.Value, nil
	}

	headerRealm := c.Request().Header.Get("Realm")
	if headerRealm != "" {
		return headerRealm, nil
	}

	return "", err
}
