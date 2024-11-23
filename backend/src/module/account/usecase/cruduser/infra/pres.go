package infra

import (
	"src/module/account/schema"
)

type ListOutput struct {
	ID         uint     `json:"id"`
	Email      string   `json:"email"`
	Mobile     *string  `json:"mobile"`
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	Admin      bool     `json:"admin"`
	RoleLabels []string `json:"role_labels"`
}

type DetailOutput struct {
	ID           uint    `json:"id"`
	Email        string  `json:"email"`
	Mobile       *string `json:"mobile"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Admin        bool    `json:"admin"`
	Locked       bool    `json:"locked"`
	LockedReason string  `json:"locked_reason"`
	RoleIDs      []uint  `json:"role_ids"`
}

func ListPres(items []schema.User) []ListOutput {
	var result []ListOutput
	for _, item := range items {
		var roleLabels []string
		for _, role := range item.Roles {
			roleLabels = append(roleLabels, role.Title)
		}
		result = append(result, ListOutput{
			ID:         item.ID,
			Email:      item.Email,
			Mobile:     item.Mobile,
			FirstName:  item.FirstName,
			LastName:   item.LastName,
			Admin:      item.Admin,
			RoleLabels: roleLabels,
		})
	}
	return result

}

func DetailPres(item schema.User) DetailOutput {
	var roleIDs []uint
	for _, role := range item.Roles {
		roleIDs = append(roleIDs, role.ID)
	}
	locked := false
	if !item.LockedAt.IsZero() {
		locked = true
	}
	return DetailOutput{
		ID:           item.ID,
		Email:        item.Email,
		Mobile:       item.Mobile,
		FirstName:    item.FirstName,
		LastName:     item.LastName,
		Admin:        item.Admin,
		Locked:       locked,
		LockedReason: item.LockedReason,
		RoleIDs:      roleIDs,
	}
}
