package infra

import (
	"src/common/ctype"
	"src/module/account/repo/iam"
	"src/module/account/repo/tenant"
	"src/module/account/repo/user"

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
		Realm:        tenant.AuthClient.Partition,
		ClientID:     tenant.AuthClient.Uid,
		ClientSecret: tenant.AuthClient.Secret,
	}, nil
}

func (r Repo) GetAuthUrl(realm string, clientId string, state ctype.Dict) string {
	iamRepo := iam.New(r.iamClient)
	return iamRepo.GetAuthUrl(realm, clientId, state)
}

func (r Repo) GetLogoutUrl(realm string, clientId string) string {
	iamRepo := iam.New(r.iamClient)
	return iamRepo.GetLogoutUrl(realm, clientId)
}

func (r Repo) CheckUserByEmail(email string) error {
	repo := user.New(r.dbClient)
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"email": email},
	}
	_, err := repo.Retrieve(queryOptions)
	return err
}

func (r Repo) CreateUser(data ctype.Dict) error {
	return nil
}
