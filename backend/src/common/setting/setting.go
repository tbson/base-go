package setting

import (
	"app/util/framework_util"
)

var DB_HOST string = framework_util.GetEnv("DB_HOST", "")
var DB_PORT string = framework_util.GetEnv("DB_PORT", "5432")
var DB_USER string = framework_util.GetEnv("DB_USER", "")
var DB_NAME string = framework_util.GetEnv("DB_NAME", "")
var DB_PASSWORD string = framework_util.GetEnv("DB_PASSWORD", "")
var TIME_ZONE string = framework_util.GetEnv("TIME_ZONE", "")
