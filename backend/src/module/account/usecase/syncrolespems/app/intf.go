package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type SynRolesPemsRepo interface {
	WritePems(pemMap ctype.PemMap) error
	EnsureRolesPems(pemMap ctype.PemMap) error
}

type RoleRepo interface {
	EnsureTenantRoles(ID uint, Uid string) error
}

type TenantRepo interface {
	List(ctype.QueryOptions) ([]schema.Tenant, error)
}