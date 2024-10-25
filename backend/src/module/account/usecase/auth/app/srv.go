package app

import (
	"context"
	"src/common/ctype"
	"src/module/account/usecase/auth/app/intf"
)

type Service struct {
	repo intf.AuthRepo
}

func (s Service) New(repo intf.AuthRepo) Service {
	return Service{repo}
}

func (srv Service) BuildAuthUrl(tenantUid string) (string, error) {
	state := ctype.Dict{
		"tenantUid": tenantUid,
	}
	authClientInfo, err := srv.repo.GetAuthClientFromTenantUid(tenantUid)
	if err != nil {
		return "", err
	}

	realm := authClientInfo.Realm
	clientId := authClientInfo.ClientID

	url := srv.repo.GetAuthUrl(realm, clientId, state)

	return url, nil
}

func (srv Service) BuildLogoutUrl(tenantUid string) (string, error) {
	authClientInfo, err := srv.repo.GetAuthClientFromTenantUid(tenantUid)
	if err != nil {
		return "", err
	}

	realm := authClientInfo.Realm
	clientId := authClientInfo.ClientID

	url := srv.repo.GetLogoutUrl(realm, clientId)

	return url, nil
}

func (srv Service) HandleCallback(
	ctx context.Context,
	tenantUid string,
	code string,
) (string, error) {
	// get realm, clientId, clientSecret from tenantUid
	// validate the callback
	// check user exist by email
	// if user not exist, create user
	// return formated response
	return "", nil
}
