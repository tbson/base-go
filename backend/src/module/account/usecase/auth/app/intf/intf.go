package intf

import (
	"context"
	"src/common/ctype"
	"src/util/ssoutil"
)

type MapRolesPems map[string][]string

type AuthUserResult struct {
	ID    uint
	Admin bool
}

type AuthClientInfo struct {
	TenantID     uint
	Realm        string
	ClientID     string
	ClientSecret string
}

type AuthRepo interface {
	GetTenantUser(tenantID uint, email string) (AuthUserResult, error)
	CreateUser(data ctype.Dict) (AuthUserResult, error)
	GetAuthClientFromTenantUid(tenantUid string) (AuthClientInfo, error)
	GetAuthUrl(realm string, clientId string, state ctype.Dict) string
	GetLogoutUrl(realm string, clientId string) string
	ValidateCallback(
		ctx context.Context,
		realm string,
		clientId string,
		clientSecret string,
		code string,
	) (ssoutil.TokensAndClaims, error)
}
