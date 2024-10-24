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
	client *gorm.DB
}

func New(client *gorm.DB) Repo {
	parent := authclient.New(client)
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
