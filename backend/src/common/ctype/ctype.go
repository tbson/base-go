package ctype

type Dict map[string]interface{}

type Role struct {
	ProfileTypes []string
	Title        string
	Module       string
	Action       string
}

type RoleMap map[string]Role

type QueryOptions struct {
	Filters  Dict
	Preloads []string
}
