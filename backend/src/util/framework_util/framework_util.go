package framework_util

import (
	"os"
)

func GetEnv(key string, defaultVal string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return defaultVal
}
