package infra

import (
	"src/module/config/repo/variable"
	"src/module/config/schema"
	"src/util/repoutil"
	"src/util/restlistutil"

	"gorm.io/gorm"
)

type Schema = schema.Variable

type Repo struct {
	*variable.Repo
	db *gorm.DB
}

func (r Repo) New(db *gorm.DB) Repo {
	parent := variable.Repo{}.New(db)
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
