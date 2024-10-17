package infra

import (
	"src/module/config/repo/common"
	"src/module/config/repo/variable"
	"src/module/config/schema"
	"src/util/restlistutil"

	"gorm.io/gorm"
)

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
) (restlistutil.ListRestfulResult[schema.Variable], error) {
	commonRepo := common.Repo[schema.Variable]{}.New(r.db)
	return commonRepo.ListPaging(options, searchableFields)
}
