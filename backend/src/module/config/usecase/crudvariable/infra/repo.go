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
	client *gorm.DB
}

func New(client *gorm.DB) Repo {
	parent := variable.New(client)
	return Repo{
		Repo:   &parent,
		client: client,
	}
}

func (r Repo) List(
	options restlistutil.ListOptions,
	searchableFields []string,
) (restlistutil.ListRestfulResult[Schema], error) {
	commonRepo := repoutil.Repo[Schema]{}.New(r.client)
	return commonRepo.ListPaging(options, searchableFields)
}
