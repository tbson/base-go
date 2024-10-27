package infra

import (
	"context"
	"src/common/ctype"
	"src/module/account/repo/iam"
	"src/module/account/repo/tenant"
	"src/module/account/repo/user"
	"src/util/ssoutil"

	"github.com/Nerzal/gocloak/v13"

	"src/module/account/usecase/auth/app/intf"

	"gorm.io/gorm"
)

type Repo struct {
	dbClient  *gorm.DB
	iamClient *gocloak.GoCloak
}

func New(dbClient *gorm.DB, iamClient *gocloak.GoCloak) Repo {
	return Repo{
		dbClient:  dbClient,
		iamClient: iamClient,
	}
}

func (r Repo) GetAuthClientFromTenantUid(tenantUid string) (intf.AuthClientInfo, error) {
	repo := tenant.New(r.dbClient)
	queryOptions := ctype.QueryOptions{
		Filters:  ctype.Dict{"uid": tenantUid},
		Preloads: []string{"AuthClient"},
	}
	tenant, err := repo.Retrieve(queryOptions)
	if err != nil {
		return intf.AuthClientInfo{}, err
	}

	return intf.AuthClientInfo{
		TenantID:     tenant.ID,
		Realm:        tenant.AuthClient.Partition,
		ClientID:     tenant.AuthClient.Uid,
		ClientSecret: tenant.AuthClient.Secret,
	}, nil
}

func (r Repo) GetMapRolesPems() (intf.MapRolesPems, error) {
	return nil, nil
}

func (r Repo) GetAuthUrl(realm string, clientId string, state ctype.Dict) string {
	iamRepo := iam.New(r.iamClient)
	return iamRepo.GetAuthUrl(realm, clientId, state)
}

func (r Repo) GetLogoutUrl(realm string, clientId string) string {
	iamRepo := iam.New(r.iamClient)
	return iamRepo.GetLogoutUrl(realm, clientId)
}

func (r Repo) GetTenantUser(
	tenantID uint,
	email string,
) (intf.AuthUserResult, error) {
	repo := user.New(r.dbClient)
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{
			"tenant_id": tenantID,
			"email":     email,
		},
	}
	user, err := repo.Retrieve(queryOptions)
	result := intf.AuthUserResult{
		ID:    user.ID,
		Admin: user.Admin,
	}
	return result, err
}

func (r Repo) CreateUser(data ctype.Dict) (intf.AuthUserResult, error) {
	repo := user.New(r.dbClient)
	user, err := repo.Create(data)

	if err != nil {
		return intf.AuthUserResult{}, err
	}

	result := intf.AuthUserResult{
		ID:    user.ID,
		Admin: user.Admin,
	}
	return result, err
}

func (r Repo) ValidateCallback(
	ctx context.Context,
	realm string,
	clientId string,
	clientSecret string,
	code string,
) (ssoutil.TokensAndClaims, error) {
	iamRepo := iam.New(r.iamClient)
	return iamRepo.ValidateCallback(ctx, realm, clientId, clientSecret, code)
}
