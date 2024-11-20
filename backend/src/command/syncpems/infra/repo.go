package infra

import (
	"slices"
	"src/common/ctype"
	"src/common/profiletype"
	"src/common/setting"
	"src/module/account/repo/pem"
	"src/module/account/repo/role"
	"src/module/account/repo/tenant"
	"src/module/account/schema"

	"gorm.io/gorm"
)

type Repo struct {
	client *gorm.DB
}

func New(client *gorm.DB) Repo {
	return Repo{
		client: client,
	}
}

func getAdmin(profileTypes []string) bool {
	// if profileTypes ONLY contains ADMIN, return true
	if len(profileTypes) == 1 && profileTypes[0] == profiletype.ADMIN {
		return true
	}

	// if profileTypes ONLY contains STAFF, return true
	if len(profileTypes) == 1 && profileTypes[0] == profiletype.STAFF {
		return true
	}

	// if profileTypes contains ADMIN and STAFF, return true
	if slices.Contains(
		profileTypes,
		profiletype.ADMIN,
	) && slices.Contains(
		profileTypes,
		profiletype.STAFF,
	) {
		return true
	}

	return false
}

func (r Repo) WritePems(pemMap ctype.PemMap) error {
	pemRepo := pem.New(r.client)
	for _, pemData := range pemMap {
		filterOptions := ctype.QueryOptions{
			Filters: ctype.Dict{
				"module": pemData.Module,
				"action": pemData.Action,
			},
		}
		data := ctype.Dict{
			"Title":  pemData.Title,
			"Module": pemData.Module,
			"Action": pemData.Action,
			"Admin":  getAdmin(pemData.ProfileTypes),
		}

		_, err := pemRepo.GetOrCreate(filterOptions, data)

		if err != nil {
			panic(err)
		}
	}
	return nil
}

func (r Repo) EnsureTenantsRoles() error {
	tenantRepo := tenant.New(r.client)
	roleRepo := role.New(r.client)

	tenants, err := tenantRepo.List(ctype.QueryOptions{})
	if err != nil {
		return err
	}

	for _, tenant := range tenants {
		profileTypes := []string{}
		if tenant.Uid == setting.ADMIN_TEANT_UID {
			profileTypes = profiletype.PlatformProfileTypes
		} else {
			profileTypes = profiletype.TenantProfileTypes
		}

		for _, roleTitle := range profileTypes {
			filterOptions := ctype.QueryOptions{
				Filters: ctype.Dict{
					"tenant_id": tenant.ID,
					"title":     roleTitle,
				},
			}
			data := ctype.Dict{
				"TenantID": tenant.ID,
				"Title":    roleTitle,
			}
			roleRepo.GetOrCreate(filterOptions, data)
		}
	}

	return nil
}

func (r Repo) EnsureRolesPems(pemMap ctype.PemMap) error {
	// get all roles
	roleRepo := role.New(r.client)
	pemRepo := pem.New(r.client)
	roles, err := roleRepo.List(ctype.QueryOptions{})
	if err != nil {
		return err
	}

	for _, role := range roles {
		newPems := []*schema.Pem{}
		// clear all pems
		r.client.Model(&role).Association("Pems").Clear()
		for _, pemData := range pemMap {
			filterOptions := ctype.QueryOptions{
				Filters: ctype.Dict{
					"module": pemData.Module,
					"action": pemData.Action,
				},
			}
			pem, err := pemRepo.Retrieve(filterOptions)
			if err != nil {
				return err
			}
			if slices.Contains(pemData.ProfileTypes, role.Title) {
				newPems = append(newPems, pem)
			}
		}
		r.client.Model(&role).Association("Pems").Append(newPems)
	}

	return nil
}
