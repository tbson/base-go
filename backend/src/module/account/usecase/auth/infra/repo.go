package infra

import (
	"src/common/ctype"
	"src/module/account/repo/tenant"
	"src/module/account/schema"

	"gorm.io/gorm"
)

type Schema = schema.Tenant

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repo {
	return Repo{
		db: db,
	}
}

func (r Repo) BuildAuthUrl(tenantUid string) (string, error) {
	repo := tenant.New(r.db)
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
