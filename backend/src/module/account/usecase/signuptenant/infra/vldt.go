package infra

type InputData struct {
	Uid        string `json:"uid" form:"uid" validate:"required"`
	Title      string `json:"title" form:"title" validate:"required"`
	AdminEmail string `json:"admin_email" form:"admin_email" validate:"required"`
}
