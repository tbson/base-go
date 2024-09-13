package routeutil

import (
	"app/common/ctype"
	"reflect"
	"runtime"
	"strings"

	"github.com/labstack/echo/v4"
)

type RuteHandlerFunc func(string, string, echo.HandlerFunc, []string, string) ctype.RoleMap

func getFnPath(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}

func getFnInfo(fnPath string) (string, string) {
	arrResult := strings.Split(fnPath, ".")
	return arrResult[0], arrResult[1]
}

func RegisterRoute(group *echo.Group, roleMap ctype.RoleMap) RuteHandlerFunc {
	return func(verb string, path string, ctrl echo.HandlerFunc, profileTypes []string, title string) ctype.RoleMap {
		verbs := []string{verb}
		group.Match(verbs, path, ctrl)
		key := getFnPath(ctrl)
		module, action := getFnInfo(key)
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
