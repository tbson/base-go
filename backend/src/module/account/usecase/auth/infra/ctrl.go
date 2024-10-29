package infra

import (
	"net/http"

	"src/common/ctype"
	"src/module/account/repo/iam"
	"src/module/account/usecase/auth/app"
	"src/util/cookieutil"
	"src/util/dbutil"
	"src/util/ssoutil"

	"github.com/labstack/echo/v4"
)

func CheckAuthUrl(c echo.Context) error {
	tenantUid := c.Param("tenantUid")
	dbClient := dbutil.Db()
	ssoClient := ssoutil.Client()
	repo := New(dbClient, ssoClient)

	srv := app.Service{}.New(repo)

	_, error := srv.GetAuthUrl(tenantUid)
	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.JSON(http.StatusOK, ctype.Dict{})
}

func GetAuthUrl(c echo.Context) error {
	tenantUid := c.Param("tenantUid")
	dbClient := dbutil.Db()
	ssoClient := ssoutil.Client()
	repo := New(dbClient, ssoClient)

	srv := app.Service{}.New(repo)

	url, error := srv.GetAuthUrl(tenantUid)
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

	url, error := srv.GetLogoutUrl(tenantUid)
	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func Callback(c echo.Context) error {
	code := c.QueryParam("code")
	state := c.QueryParam("state")
	dbClient := dbutil.Db()
	ssoClient := ssoutil.Client()
	repo := New(dbClient, ssoClient)
	srv := app.Service{}.New(repo)

	result, err := srv.HandleCallback(c.Request().Context(), state, code)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return CallbackPres(c, result)
}

func PostLogout(c echo.Context) error {
	data := map[string]interface{}{}
	return c.Render(http.StatusOK, "post-logout.html", data)
}

func RefreshToken(c echo.Context) error {
	refreshToken := cookieutil.GetValue(c, "refresh_token")
	realm := cookieutil.GetValue(c, "realm")

	iamRepo := iam.New(ssoutil.Client())
	result, err := iamRepo.RefreshToken(c.Request().Context(), realm, refreshToken)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return RefreshTokenPres(c, result)
}

func RefreshTokenCheck(c echo.Context) error {
	result := ctype.Dict{}
	return c.JSON(http.StatusOK, result)
}
