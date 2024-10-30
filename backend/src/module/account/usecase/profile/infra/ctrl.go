package infra

import (
	"net/http"
	"src/common/ctype"
	"src/module/account/repo/user"
	"src/util/dbutil"

	"github.com/labstack/echo/v4"
)

func GetProfile(c echo.Context) error {
	userID := c.Get("userID").(uint)
	client := dbutil.Db()
	userRepo := user.New(client)
	user, err := userRepo.Retrieve(ctype.QueryOptions{
		Filters: ctype.Dict{"id": userID},
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateProfile(c echo.Context) error {
	// get user id from context
	// retrieve user from db
	// get information to connect Keycloak
	// update user
	// update Keycloak user
	// need a transaction here
	return c.JSON(http.StatusOK, ctype.Dict{})
}

func ChangePassword(c echo.Context) error {
	return c.JSON(http.StatusOK, ctype.Dict{})
}
