package infra

import (
	"net/http"
	"src/common/ctype"
	"src/util/dbutil"
	"src/util/iterutil"
	"src/util/restlistutil"
	"src/util/vldtutil"

	"src/module/abstract/repo/paging"
	"src/module/account/repo/user"
	"src/module/account/schema"

	"github.com/labstack/echo/v4"
)

type Schema = schema.User

var NewRepo = user.New

var searchableFields = []string{"uid", "description", "partition"}
var filterableFields = []string{}
var orderableFields = []string{"id", "uid"}

func List(c echo.Context) error {
	pager := paging.New[Schema](dbutil.Db())

	options := restlistutil.GetOptions(c, filterableFields, orderableFields)
	listResult, error := pager.Paging(options, searchableFields)
	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.JSON(http.StatusOK, listResult)
}

func Retrieve(c echo.Context) error {
	cruder := NewRepo(dbutil.Db())

	id := vldtutil.ValidateId(c.Param("id"))
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"id": id},
	}

	result, error := cruder.Retrieve(queryOptions)

	if error != nil {
		return c.JSON(http.StatusNotFound, error)
	}

	return c.JSON(http.StatusOK, result)
}

func Create(c echo.Context) error {
	cruder := NewRepo(dbutil.Db())
	data, error := vldtutil.ValidatePayload(c, InputData{})
	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}
	result, error := cruder.Create(iterutil.StructToDict(data))
	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.JSON(http.StatusCreated, result)

}

func Update(c echo.Context) error {
	cruder := NewRepo(dbutil.Db())

	data, error := vldtutil.ValidateUpdatePayload(c, InputData{})
	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}
	id := vldtutil.ValidateId(c.Param("id"))
	result, error := cruder.Update(id, data)

	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.JSON(http.StatusOK, result)
}

func Delete(c echo.Context) error {
	cruder := NewRepo(dbutil.Db())

	id := vldtutil.ValidateId(c.Param("id"))
	ids, error := cruder.Delete(id)

	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.JSON(http.StatusOK, ids)
}

func DeleteList(c echo.Context) error {
	cruder := NewRepo(dbutil.Db())

	ids := vldtutil.ValidateIds(c.QueryParam("ids"))
	ids, error := cruder.DeleteList(ids)

	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.JSON(http.StatusOK, ids)
}
