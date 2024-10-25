package main

import (
	"src/common/ctype"
	"src/common/setting"
	"src/module/account/repo/authclient"
	"src/module/account/repo/tenant"
	"src/module/account/repo/user"
	"src/util/dbutil"
)

func main() {
	dbutil.InitDb()
	db := dbutil.Db()
	authClientRepo := authclient.Repo{}.New(db)
	tenantRepo := tenant.Repo{}.New(db)
	userRepo := user.Repo{}.New(db)

	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{
			"uid": setting.KEYCLOAK_DEFAULT_CLIENT_ID,
		},
	}
	authClientData := ctype.Dict{
		"uid":         setting.KEYCLOAK_DEFAULT_CLIENT_ID,
		"description": "Default client",
		"secret":      setting.KEYCLOAK_DEFAULT_CLIENT_SECRET,
		"partition":   setting.KEYCLOAK_DEFAULT_REALM,
		"default":     true,
	}

	authClient, err := authClientRepo.GetOrCreate(queryOptions, authClientData)
	if err != nil {
		panic(err)
	}

	queryOptions = ctype.QueryOptions{
		Filters: ctype.Dict{
			"uid": "default",
		},
	}
	tenantData := ctype.Dict{
		"auth_client_id": authClient.ID,
		"uid":            "default",
		"title":          "Default",
	}
	tenant, err := tenantRepo.GetOrCreate(queryOptions, tenantData)
	if err != nil {
		panic(err)
	}

	queryOptions = ctype.QueryOptions{
		Filters: ctype.Dict{
			"email": "admin@localhost",
		},
	}
	userData := ctype.Dict{
		"tenant_id":  tenant.ID,
		"uid":        "admin@localhost",
		"email":      "admin@localhost",
		"first_name": "Admin",
		"last_name":  "Admin",
		"admin":      true,
	}
	_, err = userRepo.GetOrCreate(queryOptions, userData)
	if err != nil {
		panic(err)
	}

}
