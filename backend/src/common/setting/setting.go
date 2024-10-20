package setting

import (
	"src/util/frameworkutil"
)

var DB_HOST string = frameworkutil.GetEnv("DB_HOST", "")
var DB_PORT string = frameworkutil.GetEnv("DB_PORT", "5432")
var DB_USER string = frameworkutil.GetEnv("DB_USER", "")
var DB_NAME string = frameworkutil.GetEnv("DB_NAME", "")
var DB_PASSWORD string = frameworkutil.GetEnv("DB_PASSWORD", "")
var TIME_ZONE string = frameworkutil.GetEnv("TIME_ZONE", "")

var KEYCLOAK_URL string = frameworkutil.GetEnv("KEYCLOAK_URL", "")
var KEYCLOAK_API_URL string = frameworkutil.GetEnv("KEYCLOAK_API_URL", "")
var KEYCLOAK_DEFAULT_REALM string = frameworkutil.GetEnv("KEYCLOAK_DEFAULT_REALM", "")
var KEYCLOAK_DEFAULT_CLIENT_ID string = frameworkutil.GetEnv("KEYCLOAK_DEFAULT_CLIENT_ID", "")
var KEYCLOAK_DEFAULT_CLIENT_SECRET string = frameworkutil.GetEnv("KEYCLOAK_DEFAULT_CLIENT_SECRET", "")
var KEYCLOAK_DEFAULT_REDIRECT_URI string = frameworkutil.GetEnv("KEYCLOAK_DEFAULT_REDIRECT_URI", "")

const DEFAULT_LANG = "en"
