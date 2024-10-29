package app

import (
	"context"
	"src/common/ctype"
	"src/util/errutil"
	"src/util/iterutil"
	"src/util/localeutil"
	"src/util/ssoutil"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Service struct {
	repo AuthRepo
}

func (s Service) New(repo AuthRepo) Service {
	return Service{repo}
}

func (srv Service) parseTenantUidFromState(state string) (string, error) {
	localizer := localeutil.Get()
	stateData, err := ssoutil.DecodeState(state)
	if err != nil {
		return "", err
	}

	tenantUid, ok := stateData["tenantUid"].(string)
	if !ok {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.InvalidState,
		})
		return "", errutil.New("", []string{msg})
	}
	return tenantUid, nil
}

func (srv Service) GetAuthUrl(tenantUid string) (string, error) {
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

func (srv Service) GetLogoutUrl(tenantUid string) (string, error) {
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
	state string,
	code string,
) (ssoutil.TokensAndClaims, error) {
	tenantUid, err := srv.parseTenantUidFromState(state)
	if err != nil {
		return ssoutil.TokensAndClaims{}, err
	}

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

	user, err := srv.repo.GetTenantUser(tenantID, userInfo.Email)

	if err != nil {
		userData := iterutil.StructToDict(userInfo)
		userData["TenantID"] = tenantID
		_, err = srv.repo.CreateUser(userData)
		if err != nil {
			return tokensAndClaims, err
		}
	}

	if user.Admin {
		tokensAndClaims.UserInfo.ProfileType = "admin"
	} else {
		tokensAndClaims.UserInfo.ProfileType = "user"
	}
	tokensAndClaims.UserInfo.TenantUid = tenantUid
	tokensAndClaims.UserInfo.ID = user.ID
	return tokensAndClaims, nil
}

func (srv Service) GetPemModulesActionsMap(
	userId uint,
) (PemModulesActionsMap, error) {
	return srv.repo.GetPemModulesActionsMap(userId)
}
