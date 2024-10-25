package infra

import (
	"net/http"

	"src/common/setting"
	"src/module/account/repo/iam"
	"src/util/dbutil"
	"src/util/errutil"
	"src/util/localeutil"
	"src/util/ssoutil"

	"src/module/account/usecase/auth/app"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
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
	localizer := localeutil.Get()
	code := c.QueryParam("code")
	state := c.QueryParam("state")
	dbClient := dbutil.Db()
	ssoClient := ssoutil.Client()
	repo := New(dbClient, ssoClient)
	srv := app.Service{}.New(repo)

	stateData, err := ssoutil.DecodeState(state)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	tenantUid, ok := stateData["tenantUid"].(string)
	if !ok {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.InvalidState,
		})
		return c.JSON(http.StatusBadRequest, errutil.New("", []string{msg}))
	}

	result, err := srv.HandleCallback(c.Request().Context(), tenantUid, code)
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
