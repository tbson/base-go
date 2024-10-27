package routeutil

import (
	"reflect"
	"runtime"
	"src/common/ctype"
	"strings"

	"github.com/labstack/echo/v4"
)

type RuteHandlerFunc func(string, string, echo.HandlerFunc, []string, string) ctype.RoleMap

func getFnPath(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}

func getFnInfo(fnPath string) (string, string) {
	arrResult := strings.Split(fnPath, ".")
	module := arrResult[0]
	action := arrResult[1]

	arrModule := strings.Split(module, "/")
	module = arrModule[len(arrModule)-2]
	return module, action
}

func RegisterRoute(group *echo.Group, roleMap ctype.RoleMap) RuteHandlerFunc {
	return func(
		verb string,
		path string,
		ctrl echo.HandlerFunc,
		profileTypes []string,
		title string,
	) ctype.RoleMap {
		verbs := []string{verb}
		group.Match(verbs, path, ctrl)
		if len(profileTypes) == 0 || len(title) == 0 {
			return roleMap
		}
		fnPath := getFnPath(ctrl)
		module, action := getFnInfo(fnPath)
		key := module + "." + action
		role := ctype.Role{
			ProfileTypes: profileTypes,
			Title:        title,
			Module:       module,
			Action:       action,
		}
		roleMap[key] = role
		return roleMap
	}
}
