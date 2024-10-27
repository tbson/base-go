package routeutil

import (
	"reflect"
	"runtime"
	"src/common/ctype"
	"src/middleware"
	"strings"

	"github.com/labstack/echo/v4"
)

type RuteHandlerFunc func(string, string, echo.HandlerFunc, []string, string) ctype.PemMap

func getFnPath(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}

func GetHandlerInfo(ctrl echo.HandlerFunc) (string, string) {
	fnPath := getFnPath(ctrl)
	arrResult := strings.Split(fnPath, ".")
	module := arrResult[0]
	action := arrResult[1]

	arrModule := strings.Split(module, "/")
	module = arrModule[len(arrModule)-2]
	return module, action
}

func RegisterRoute(group *echo.Group, pemMap ctype.PemMap) RuteHandlerFunc {
	return func(
		verb string,
		path string,
		ctrl echo.HandlerFunc,
		profileTypes []string,
		title string,
	) ctype.PemMap {
		verbs := []string{verb}
		module, action := GetHandlerInfo(ctrl)
		if len(profileTypes) > 0 && len(title) > 0 {
			key := module + "." + action
			role := ctype.Pem{
				ProfileTypes: profileTypes,
				Title:        title,
				Module:       module,
				Action:       action,
			}
			pemMap[key] = role
			group.Match(verbs, path, ctrl, middleware.AuthMiddleware(module, action))
		} else {
			group.Match(verbs, path, ctrl)
		}
		return pemMap
	}
}
