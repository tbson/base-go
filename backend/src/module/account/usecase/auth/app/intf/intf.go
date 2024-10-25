package intf

import (
	"src/common/ctype"
)

type AuthClientInfo struct {
	Realm        string
	ClientID     string
	ClientSecret string
}

type AuthRepo interface {
	CheckUserByEmail(email string) error
	CreateUser(data ctype.Dict) error
	GetAuthClientFromTenantUid(tenantUid string) (AuthClientInfo, error)
	GetAuthUrl(realm string, clientId string, state ctype.Dict) string
	GetLogoutUrl(realm string, clientId string) string
}
