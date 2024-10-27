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

func (r Repo) WritePems(roleMap ctype.RoleMap) error {
	pemRepo := pem.New(r.client)
	for _, pemData := range roleMap {
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

func (r Repo) EnsureRolesPems(roleMap ctype.RoleMap) error {
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
		for _, pemInfo := range roleMap {
			pemData := ctype.QueryOptions{
				Filters: ctype.Dict{
					"module": pemInfo.Module,
					"action": pemInfo.Action,
				},
			}
			pem, err := pemRepo.Retrieve(pemData)
			if err != nil {
				return err
			}
			if slices.Contains(pemInfo.ProfileTypes, role.Title) {
				newPems = append(newPems, pem)
			}
		}
		r.client.Model(&role).Association("Pems").Append(newPems)
	}

	return nil
}
