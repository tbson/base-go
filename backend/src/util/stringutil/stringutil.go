package stringutil

func ToCapital(str string) string {
	if len(str) == 0 {
		return str
	}
	return str[:1] + str[1:]
}
