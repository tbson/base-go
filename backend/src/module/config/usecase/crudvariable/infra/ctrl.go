package infra

import (
	"net/http"
	"src/common/ctype"
	"src/util/dbutil"
	"src/util/iterutil"
	"src/util/restlistutil"
	"src/util/vldtutil"

	"src/module/abstract/repo/paging"
	"src/module/config/repo/variable"
	"src/module/config/schema"

	"github.com/labstack/echo/v4"
)

var searchableFields = []string{"key", "value", "description"}
var filterableFields = []string{"data_type"}
var orderableFields = []string{"id", "key"}

func List(c echo.Context) error {
	pager := paging.New[schema.Variable](dbutil.Db())

	options := restlistutil.GetOptions(c, filterableFields, orderableFields)
	listResult, error := pager.Paging(options, searchableFields)
	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.JSON(http.StatusOK, listResult)
}

func Retrieve(c echo.Context) error {
	cruder := variable.New(dbutil.Db())

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
	cruder := variable.New(dbutil.Db())
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
	cruder := variable.New(dbutil.Db())

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
	cruder := variable.New(dbutil.Db())

	id := vldtutil.ValidateId(c.Param("id"))
	ids, error := cruder.Delete(id)

	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.JSON(http.StatusOK, ids)
}

func DeleteList(c echo.Context) error {
	cruder := variable.New(dbutil.Db())

	ids := vldtutil.ValidateIds(c.QueryParam("ids"))
	ids, error := cruder.DeleteList(ids)

	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.JSON(http.StatusOK, ids)
}
