package infra

import (
	"src/module/account/repo/tenant"
	"src/module/account/schema"
	"src/util/repoutil"
	"src/util/restlistutil"

	"gorm.io/gorm"
)

type Schema = schema.Tenant

type Repo struct {
	*tenant.Repo
	db *gorm.DB
}

func (r Repo) New(db *gorm.DB) Repo {
	parent := tenant.Repo{}.New(db)
	return Repo{
		Repo: &parent,
		db:   db,
	}
}

func (r Repo) List(
	options restlistutil.ListOptions,
	searchableFields []string,
) (restlistutil.ListRestfulResult[Schema], error) {
	commonRepo := repoutil.Repo[Schema]{}.New(r.db)
	return commonRepo.ListPaging(options, searchableFields)
}
