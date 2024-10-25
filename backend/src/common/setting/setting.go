package setting

import (
	"fmt"
	"src/util/frameworkutil"
)

var BASE_URL string = frameworkutil.GetEnv("BASE_URL", "")
var DOMAIN string = frameworkutil.GetEnv("DOMAIN", "")
var DB_HOST string = frameworkutil.GetEnv("DB_HOST", "")
var DB_PORT string = frameworkutil.GetEnv("DB_PORT", "5432")
var DB_USER string = frameworkutil.GetEnv("DB_USER", "")
var DB_NAME string = frameworkutil.GetEnv("DB_NAME", "")
var DB_PASSWORD string = frameworkutil.GetEnv("DB_PASSWORD", "")
var TIME_ZONE string = frameworkutil.GetEnv("TIME_ZONE", "")

var KEYCLOAK_URL string = frameworkutil.GetEnv("KEYCLOAK_URL", "")
var KEYCLOAK_DEFAULT_REALM string = frameworkutil.GetEnv("KEYCLOAK_DEFAULT_REALM", "")
var KEYCLOAK_DEFAULT_CLIENT_ID string = frameworkutil.GetEnv("KEYCLOAK_DEFAULT_CLIENT_ID", "")
var KEYCLOAK_DEFAULT_CLIENT_SECRET string = frameworkutil.GetEnv("KEYCLOAK_DEFAULT_CLIENT_SECRET", "")
var KEYCLOAK_DEFAULT_REDIRECT_URI string = fmt.Sprintf(
	"%s%s",
	BASE_URL,
	frameworkutil.GetEnv("KEYCLOAK_DEFAULT_REDIRECT_URI", ""),
)
var KEYCLOAK_DEFAULT_POST_LOGOUT_URI string = fmt.Sprintf(
	"%s%s",
	BASE_URL,
	frameworkutil.GetEnv("KEYCLOAK_DEFAULT_POST_LOGOUT_URI", ""),
)

const DEFAULT_LANG = "en"
