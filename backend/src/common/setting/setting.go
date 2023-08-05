package setting

import (
	"app/util/framework_util"
)

var DB_HOST string = framework_util.GetEnv("DB_HOST", "")
var DB_NAME string = framework_util.GetEnv("DB_NAME", "")
var DB_USER string = framework_util.GetEnv("DB_USER", "")
var DB_PASSWORD string = framework_util.GetEnv("DB_PASSWORD", "")
