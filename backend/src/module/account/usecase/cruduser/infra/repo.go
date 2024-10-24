package infra

import (
	"src/module/account/repo/user"
	"src/module/account/schema"
	"src/util/repoutil"
	"src/util/restlistutil"

	"gorm.io/gorm"
)

type Schema = schema.User

type Repo struct {
	*user.Repo
	db *gorm.DB
}

func New(db *gorm.DB) Repo {
	parent := user.New(db)
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
