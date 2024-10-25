package infra

import (
	"net/http"

	"src/common/setting"
	"src/module/account/repo/iam"
	"src/util/dbutil"
	"src/util/ssoutil"

	"src/module/account/usecase/auth/app"

	"github.com/labstack/echo/v4"
)

func GetAuthUrl(c echo.Context) error {
	tenantUid := c.Param("tenantUid")
	dbClient := dbutil.Db()
	ssoClient := ssoutil.Client()
	repo := New(dbClient, ssoClient)

	srv := app.Service{}.New(repo)

	url, error := srv.BuildAuthUrl(tenantUid)
	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func GetLogoutUrl(c echo.Context) error {
	tenantUid := c.Param("tenantUid")
	dbClient := dbutil.Db()
	ssoClient := ssoutil.Client()
	repo := New(dbClient, ssoClient)

	srv := app.Service{}.New(repo)

	url, error := srv.BuildLogoutUrl(tenantUid)
	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func Callback(c echo.Context) error {
	iamRepo := iam.New(ssoutil.Client())
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
	clientId := setting.KEYCLOAK_DEFAULT_CLIENT_ID
	clientSecret := setting.KEYCLOAK_DEFAULT_CLIENT_SECRET
	result, err := iamRepo.ValidateCallback(c.Request().Context(), realm, clientId, clientSecret, code)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func RefreshToken(c echo.Context) error {
	iamRepo := iam.New(ssoutil.Client())
	realm := setting.KEYCLOAK_DEFAULT_REALM
	refreshToken := c.FormValue("refresh_token")
	result, err := iamRepo.RefreshToken(c.Request().Context(), realm, refreshToken)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
