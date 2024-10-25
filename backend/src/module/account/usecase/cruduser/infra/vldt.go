package infra

import "src/common/ctype"

type InputData struct {
	TenantID    uint       `json:"tenant_id" validate:"required"`
	TenantTmpID *uint      `json:"tenant_tmp_id"`
	Uid         string     `json:"title" validate:"required"`
	Email       string     `json:"title" validate:"required"`
	Mobile      *string    `json:"title"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Avatar      string     `json:"avatar"`
	AvatarStr   string     `json:"avatar_str"`
	ExtraInfo   ctype.Dict `json:"extra_info"`
	Admin       bool       `json:"admin"`
}
