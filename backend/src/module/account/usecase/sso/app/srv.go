package app

import (
	"src/module/account/usecase/sso/app/intf"
)

type Service struct {
	repo intf.SsoRepo
}

func (s Service) New(repo intf.SsoRepo) Service {
	return Service{repo}
}

func (srv Service) BuildAuthUrl(tenantUid string) (string, error) {
	return "", nil
}

func (srv Service) BuildLogoutUrl(tenantUid string) (string, error) {
	return "", nil
}

func (srv Service) HandleCallback(tenantUid string) (string, error) {
	return "", nil
}
