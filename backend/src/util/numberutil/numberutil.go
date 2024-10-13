package numberutil

import (
	"strconv"
)

func StrToInt(id string, defaultValue int) int {
	if id == "" {
		return defaultValue
	}
	if id, err := strconv.Atoi(id); err == nil {
		return id
	}
	return defaultValue
}
