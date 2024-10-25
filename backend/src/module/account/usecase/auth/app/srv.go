package app

import (
	"context"
	"src/common/ctype"
	"src/module/account/usecase/auth/app/intf"
	"src/util/iterutil"
	"src/util/ssoutil"
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
) (ssoutil.TokensAndClaims, error) {
	authClientInfo, err := srv.repo.GetAuthClientFromTenantUid(tenantUid)
	if err != nil {
		return ssoutil.TokensAndClaims{}, err
	}
	tenantID := authClientInfo.TenantID
	realm := authClientInfo.Realm
	clientId := authClientInfo.ClientID
	clientSecret := authClientInfo.ClientSecret

	tokensAndClaims, err := srv.repo.ValidateCallback(
		ctx, realm, clientId, clientSecret, code,
	)

	if err != nil {
		return ssoutil.TokensAndClaims{}, err
	}

	userInfo := tokensAndClaims.UserInfo

	err = srv.repo.CheckUserByEmail(userInfo.Email)

	if err != nil {
		userData := iterutil.StructToDict(userInfo)
		userData["TenantID"] = tenantID
		err = srv.repo.CreateUser(userData)
		if err != nil {
			return tokensAndClaims, err
		}
	}

	return tokensAndClaims, nil
}
