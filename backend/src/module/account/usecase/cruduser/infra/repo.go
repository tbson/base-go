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
	client *gorm.DB
}

func New(client *gorm.DB) Repo {
	parent := user.New(client)
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
