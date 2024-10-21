package app

import (
	"encoding/json"
	"src/common/ctype"
	"src/common/intf"
	"src/module/account/schema"
	"src/util/restlistutil"

	"gorm.io/datatypes"
)

type Schema = schema.User

type Data struct {
	TenantID    uint       `json:"tenant_id" validate:"required"`
	TenantTmpID *uint      `json:"tenant_tmp_id"`
	Uid         string     `json:"title" validate:"required"`
	Email       string     `json:"title" validate:"required"`
	Mobile      *string    `json:"title"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Avatar      string     `json:"avatar"`
	AvatarStr   string     `json:"avatar_str"`
	ExtraInfo   ctype.Dict `json:"extra_info"`
	Admin       bool       `json:"admin"`
}

func (data Data) ToSchema() *Schema {
	extraInfoJSON, err := json.Marshal(data.ExtraInfo)
	if err != nil {
		// Handle the error (you can return an error or handle it in another way)
		panic("Failed to marshal ExtraInfo")
	}
	return &Schema{
		TenantID:    data.TenantID,
		TenantTmpID: data.TenantTmpID,
		Uid:         data.Uid,
		Email:       data.Email,
		Mobile:      data.Mobile,
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Avatar:      data.Avatar,
		AvatarStr:   data.AvatarStr,
		ExtraInfo:   datatypes.JSON(extraInfoJSON),
		Admin:       data.Admin,
	}
}

type Service struct {
	repo intf.RestCrudRepo[Schema]
}

func (s Service) New(repo intf.RestCrudRepo[Schema]) Service {
	return Service{repo}
}

func (srv Service) List(options restlistutil.ListOptions, searchableFields []string) (restlistutil.ListRestfulResult[Schema], error) {
	return srv.repo.List(options, searchableFields)
}

func (srv Service) Retrieve(queryOptions ctype.QueryOptions) (*Schema, error) {
	return srv.repo.Retrieve(queryOptions)
}

func (srv Service) Create(inputData Data) (*Schema, error) {
	schema := inputData.ToSchema()
	return srv.repo.Create(schema)
}

func (srv Service) Update(id int, inputData ctype.Dict) (*Schema, error) {
	return srv.repo.Update(id, inputData)
}

func (srv Service) Delete(id int) ([]int, error) {
	return srv.repo.Delete(id)
}

func (srv Service) DeleteList(ids []int) ([]int, error) {
	return srv.repo.DeleteList(ids)
}
