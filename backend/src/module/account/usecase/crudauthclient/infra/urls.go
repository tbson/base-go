package infra

import (
	"fmt"
	"src/util/routeutil"

	"src/common/ctype"
	"src/common/profiletype"

	"github.com/labstack/echo/v4"
)

var module = "account"
var useCaseGroup = "auth-client"
var useCaseGroupName = "auth client"

func RegisterUrls(e *echo.Group, pemMap ctype.PemMap) (*echo.Group, ctype.PemMap) {
	g := e.Group(fmt.Sprintf("/%s/%s", module, useCaseGroup))
	rr := routeutil.RegisterRoute(g, pemMap)

	rr.Rbac(
		"GET", "/", List,
		[]string{profiletype.ADMIN},
		fmt.Sprintf("Get %s list", useCaseGroupName),
	)
	rr.Rbac(
		"GET", "/:id", Retrieve,
		[]string{profiletype.ADMIN},
		fmt.Sprintf("Get %s detail", useCaseGroupName),
	)
	rr.Rbac(
		"POST", "/", Create,
		[]string{profiletype.ADMIN},
		fmt.Sprintf("Create %s", useCaseGroupName),
	)
	rr.Rbac(
		"PUT", "/:id", Update,
		[]string{profiletype.ADMIN},
		fmt.Sprintf("Update %s", useCaseGroupName),
	)
	rr.Rbac(
		"DELETE", "/:id", Delete,
		[]string{profiletype.ADMIN},
		fmt.Sprintf("Delete %s", useCaseGroupName),
	)
	rr.Rbac(
		"DELETE", "/", DeleteList,
		[]string{profiletype.ADMIN},
		fmt.Sprintf("Delete list %s", useCaseGroupName),
	)
	return e, pemMap
}
