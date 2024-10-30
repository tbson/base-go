package infra

import (
	"slices"
	"src/common/ctype"
	"src/module/account/repo/tenant"
	"src/module/account/repo/user"
	"src/util/stringutil"

	"src/module/account/usecase/auth/app"

	"gorm.io/gorm"
)

type Repo struct {
	client *gorm.DB
}

func New(client *gorm.DB) Repo {
	return Repo{
		client: client,
	}
}

func (r Repo) GetTenantUser(
	tenantID uint,
	email string,
) (app.AuthUserResult, error) {
	repo := user.New(r.client)
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{
			"tenant_id": tenantID,
			"email":     email,
		},
	}
	user, err := repo.Retrieve(queryOptions)
	if err != nil {
		return app.AuthUserResult{}, err
	}
	result := app.AuthUserResult{
		ID:    user.ID,
		Admin: user.Admin,
	}
	return result, err
}

func (r Repo) GetAuthClientFromTenantUid(tenantUid string) (app.AuthClientInfo, error) {
	repo := tenant.New(r.client)
	queryOptions := ctype.QueryOptions{
		Filters:  ctype.Dict{"uid": tenantUid},
		Preloads: []string{"AuthClient"},
	}
	tenant, err := repo.Retrieve(queryOptions)
	if err != nil {
		return app.AuthClientInfo{}, err
	}

	return app.AuthClientInfo{
		TenantID:     tenant.ID,
		Realm:        tenant.AuthClient.Partition,
		ClientID:     tenant.AuthClient.Uid,
		ClientSecret: tenant.AuthClient.Secret,
	}, nil
}

func (r Repo) GetPemModulesActionsMap(userId uint) (app.PemModulesActionsMap, error) {
	repo := user.New(r.client)

	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"id": userId},
		Preloads: []string{
			"Roles.Pems",
		},
	}
	user, err := repo.Retrieve(queryOptions)
	if err != nil {
		return nil, err
	}

	result := make(app.PemModulesActionsMap)
	for _, role := range user.Roles {
		for _, pem := range role.Pems {
			module := stringutil.ToSnakeCase(pem.Module)
			action := stringutil.ToSnakeCase(pem.Action)
			if _, ok := result[module]; !ok {
				result[module] = make([]string, 0)
			}
			if !slices.Contains(result[module], action) {
				result[module] = append(result[module], action)
			}
		}
	}

	return result, nil
}
