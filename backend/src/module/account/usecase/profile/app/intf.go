package app

import (
	"src/common/ctype"
	"src/module/account/schema"
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

type UserRepo interface {
	Retrieve(opts ctype.QueryOptions) (*schema.User, error)
	Update(id int, data ctype.Dict) (*schema.User, error)
}

type IamRepo interface {
	GetAdminAccessToken() (string, error)
	UpdateUser(accessToken string, realm string, sub string, data ctype.Dict) error
	SetPassword(accessToken string, sub string, realm string, password string) error
}
