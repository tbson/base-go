package infra

import (
	"src/common/ctype"
	"src/module/account/repo/tenant"
	"src/module/account/schema"

	"gorm.io/gorm"
)

type Schema = schema.Tenant

type Repo struct {
	client *gorm.DB
}

func New(client *gorm.DB) Repo {
	return Repo{
		client: client,
	}
}

func (r Repo) BuildAuthUrl(tenantUid string) (string, error) {
	repo := tenant.New(r.client)
	queryOptions := ctype.QueryOptions{
		Filters:  ctype.Dict{"uid": tenantUid},
		Preloads: []string{"AuthClient"},
	}
	tenant, err := repo.Retrieve(queryOptions)
	if err != nil {
		return "", err
	}

	return tenant.Uid, nil
}
