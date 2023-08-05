package ctype

type Role struct {
	ProfileTypes []string
	Title        string
	Module       string
	Action       string
}

type RoleMap map[string]Role
