package infra

import (
	"net/http"

	"src/common/ctype"
	"src/common/setting"
	"src/util/ssoutil"

	"github.com/labstack/echo/v4"
)

func GetAuthUrl(c echo.Context) error {

	// tenantUid := c.Param("tenantUid")
	state := ctype.Dict{
		"tenantId": c.Param("tenantUid"),
	}
	realm := setting.KEYCLOAK_DEFAULT_REALM
	clientId := setting.KEYCLOAK_DEFAULT_CLIENT_ID

	url := ssoutil.GetAuthUrl(realm, clientId, state)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func GetLogoutUrl(c echo.Context) error {
	realm := setting.KEYCLOAK_DEFAULT_REALM
	clientId := setting.KEYCLOAK_DEFAULT_CLIENT_ID

	url := ssoutil.GetLogoutUrl(realm, clientId)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func Callback(c echo.Context) error {
	code := c.QueryParam("code")
	// Decode the state
	/*
		stateStr := c.QueryParam("state")
		stateData, err := decodeState(stateStr)
		if err != nil {
			msg := localizer.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: localeutil.InvalidState,
			})
			return result, errutil.New("", []string{msg})
		}
		realm, ok := stateData["realm"].(string)
		if !ok {
			msg := localizer.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: localeutil.NoRealmFound,
			})
			return result, errutil.New("", []string{msg})
		}
	*/
	realm := setting.KEYCLOAK_DEFAULT_REALM
	result, err := ssoutil.ValidateCallback(c.Request().Context(), realm, code)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func RefreshToken(c echo.Context) error {
	realm := setting.KEYCLOAK_DEFAULT_REALM
	refreshToken := c.FormValue("refresh_token")
	result, err := ssoutil.RefreshToken(c.Request().Context(), realm, refreshToken)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
