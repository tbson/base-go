package infra

import (
	"src/module/account/repo/authclient"
	"src/module/account/schema"
	"src/util/repoutil"
	"src/util/restlistutil"

	"gorm.io/gorm"
)

type Schema = schema.AuthClient

type Repo struct {
	*authclient.Repo
	db *gorm.DB
}

func (r Repo) New(db *gorm.DB) Repo {
	parent := authclient.Repo{}.New(db)
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
