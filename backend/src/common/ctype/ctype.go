package ctype

type Dict map[string]interface{}

type Pem struct {
	ProfileTypes []string
	Title        string
	Module       string
	Action       string
}

type PemMap map[string]Pem

type QueryOptions struct {
	Filters  Dict
	Preloads []string
}
