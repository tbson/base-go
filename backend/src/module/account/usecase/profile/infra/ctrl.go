package infra

import (
	"context"
	"net/http"
	"src/common/ctype"
	"src/module/account/repo/iam"
	"src/module/account/repo/user"
	"src/util/dbutil"
	"src/util/ssoutil"
	"src/util/vldtutil"

	"github.com/Nerzal/gocloak/v13"
	"github.com/labstack/echo/v4"
)

type InputData struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Mobile    string `json:"mobile" validate:"required"`
}

type InputPassword struct {
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"password_confirm" validate:"required"`
}

func GetProfile(c echo.Context) error {
	userID := c.Get("userID").(uint)
	client := dbutil.Db()
	userRepo := user.New(client)
	user, err := userRepo.Retrieve(ctype.QueryOptions{
		Filters:  ctype.Dict{"id": userID},
		Preloads: []string{"Tenant.AuthClient"},
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateProfile(c echo.Context) error {
	userID := c.Get("userID").(uint)
	ssoClient := ssoutil.Client()
	iamRepo := iam.New(ssoClient)
	client := dbutil.Db()
	userRepo := user.New(client)
	user, err := userRepo.Retrieve(ctype.QueryOptions{
		Filters:  ctype.Dict{"id": userID},
		Preloads: []string{"Tenant.AuthClient"},
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	sub := user.Sub
	realm := user.Tenant.AuthClient.Partition

	data, error := vldtutil.ValidateUpdatePayload(c, InputData{})
	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}
	accessToken, err := iamRepo.GetAdminAccessToken()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	userData := gocloak.User{
		ID:        sub,
		FirstName: gocloak.StringP(data["FirstName"].(string)),
		LastName:  gocloak.StringP(data["LastName"].(string)),
		Attributes: &map[string][]string{
			"mobile": {data["Mobile"].(string)},
		},
	}
	ctx := context.Background()
	error = ssoClient.UpdateUser(ctx, accessToken, realm, userData)

	userResult, err := userRepo.Update(int(userID), data)
	return c.JSON(http.StatusOK, userResult)
}

func ChangePassword(c echo.Context) error {
	userID := c.Get("userID").(uint)
	ssoClient := ssoutil.Client()
	iamRepo := iam.New(ssoClient)
	client := dbutil.Db()
	userRepo := user.New(client)
	user, err := userRepo.Retrieve(ctype.QueryOptions{
		Filters:  ctype.Dict{"id": userID},
		Preloads: []string{"Tenant.AuthClient"},
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	sub := user.Sub
	realm := user.Tenant.AuthClient.Partition

	data, error := vldtutil.ValidateUpdatePayload(c, InputPassword{})
	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	if data["Password"].(string) != data["PasswordConfirm"].(string) {
		msg := "Password and Password Confirm do not match"
		return c.JSON(http.StatusBadRequest, msg)
	}

	accessToken, err := iamRepo.GetAdminAccessToken()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	password := data["Password"].(string)
	ctx := context.Background()
	error = ssoClient.SetPassword(ctx, accessToken, *sub, realm, password, false)

	return c.JSON(http.StatusOK, ctype.Dict{})
}
