package setting

import (
	"app/util/frameworkutil"
)

var DB_HOST string = frameworkutil.GetEnv("DB_HOST", "")
var DB_PORT string = frameworkutil.GetEnv("DB_PORT", "5432")
var DB_USER string = frameworkutil.GetEnv("DB_USER", "")
var DB_NAME string = frameworkutil.GetEnv("DB_NAME", "")
var DB_PASSWORD string = frameworkutil.GetEnv("DB_PASSWORD", "")
var TIME_ZONE string = frameworkutil.GetEnv("TIME_ZONE", "")
