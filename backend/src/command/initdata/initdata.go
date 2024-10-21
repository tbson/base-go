package main

import (
	"src/common/ctype"
	"src/common/setting"
	"src/module/account/repo/authclient"
	"src/module/account/repo/tenant"
	"src/module/account/repo/user"
	"src/module/account/schema"
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
	authClientData := schema.AuthClient{
		Uid:         setting.KEYCLOAK_DEFAULT_CLIENT_ID,
		Description: "Default client",
		Secret:      setting.KEYCLOAK_DEFAULT_CLIENT_SECRET,
		Partition:   setting.KEYCLOAK_DEFAULT_REALM,
		Default:     true,
	}

	authClient, err := authClientRepo.GetOrCreate(queryOptions, &authClientData)
	if err != nil {
		panic(err)
	}

	queryOptions = ctype.QueryOptions{
		Filters: ctype.Dict{
			"uid": "default",
		},
	}
	tenantData := schema.Tenant{
		AuthClientID: authClient.ID,
		Uid:          "default",
		Title:        "Default",
	}
	tenant, err := tenantRepo.GetOrCreate(queryOptions, &tenantData)
	if err != nil {
		panic(err)
	}

	queryOptions = ctype.QueryOptions{
		Filters: ctype.Dict{
			"email": "admin@localhost",
		},
	}
	userData := schema.User{
		TenantID:  tenant.ID,
		Uid:       "admin@localhost",
		Email:     "admin@localhost",
		FirstName: "Admin",
		LastName:  "Admin",
		Admin:     true,
	}
	_, err = userRepo.GetOrCreate(queryOptions, &userData)
	if err != nil {
		panic(err)
	}

}
