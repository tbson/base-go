package infra

import (
	"net/http"
	"src/common/ctype"

	"github.com/labstack/echo/v4"
)

func SignupTenant(c echo.Context) error {
	/*
		Asking for:
		- uid
		- title
		- admin email

		Create tenant with default auth client (uid, title)

		Seeding roles

		Create admin user (email)

		Redirect to signup page directly
	*/
	result := ctype.Dict{}
	return c.JSON(http.StatusOK, result)
}
