package infra

import (
	"net/http"
	"src/common/ctype"
	"src/util/dbutil"
	"src/util/iterutil"
	"src/util/restlistutil"
	"src/util/vldtutil"

	"src/module/account/usecase/cruduser/app"

	"github.com/labstack/echo/v4"
)

var searchableFields = []string{"uid", "description", "partition"}
var filterableFields = []string{}
var orderableFields = []string{"id", "uid"}

func List(c echo.Context) error {
	repo := New(dbutil.Db())
	srv := app.Service{}.New(repo)

	options := restlistutil.GetOptions(c, filterableFields, orderableFields)
	listResult, error := srv.List(options, searchableFields)
	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.JSON(http.StatusOK, listResult)
}

func Retrieve(c echo.Context) error {
	repo := New(dbutil.Db())
	srv := app.Service{}.New(repo)

	id := vldtutil.ValidateId(c.Param("id"))
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"id": id},
	}
	result, error := srv.Retrieve(queryOptions)

	if error != nil {
		return c.JSON(http.StatusNotFound, error)
	}

	return c.JSON(http.StatusOK, result)
}

func Create(c echo.Context) error {
	repo := New(dbutil.Db())
	srv := app.Service{}.New(repo)

	data, error := vldtutil.ValidatePayload(c, InputData{})
	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	result, error := srv.Create(iterutil.StructToDict(data))
	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.JSON(http.StatusCreated, result)

}

func Update(c echo.Context) error {
	repo := New(dbutil.Db())
	srv := app.Service{}.New(repo)

	data, error := vldtutil.ValidateUpdatePayload(c, InputData{})
	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	id := vldtutil.ValidateId(c.Param("id"))
	result, error := srv.Update(id, data)

	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.JSON(http.StatusOK, result)
}

func Delete(c echo.Context) error {
	repo := New(dbutil.Db())
	srv := app.Service{}.New(repo)

	id := vldtutil.ValidateId(c.Param("id"))
	ids, error := srv.Delete(id)

	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.JSON(http.StatusOK, ids)
}

func DeleteList(c echo.Context) error {
	repo := New(dbutil.Db())
	srv := app.Service{}.New(repo)

	ids := vldtutil.ValidateIds(c.QueryParam("ids"))
	ids, error := srv.DeleteList(ids)

	if error != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	return c.JSON(http.StatusOK, ids)
}
