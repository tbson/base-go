package infra

import (
	"src/util/errutil"
	"src/util/localeutil"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type InputData struct {
	TenantID uint   `json:"tenant_id" form:"tenant_id" validate:"required"`
	Title    string `json:"title" form:"title" validate:"required"`
	PemIDs   []uint `json:"pem_ids" form:"pem_ids" validate:"required"`
}

func CheckRequiredFilter(c echo.Context, param string) error {
	localizer := localeutil.Get()
	if c.QueryParam(param) == "" {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.MissingTenantID,
		})
		return errutil.New("", []string{msg})
	}
	return nil
}
