package schema

import (
	"encoding/json"
	"time"

	"src/common/ctype"

	"gorm.io/datatypes"
)

type AuthClient struct {
	ID          uint
	Tenants     []Tenant `gorm:"constraint:OnDelete:SET NULL;"`
	Uid         string   `gorm:"type:text;not null;unique"`
	Description string   `gorm:"type:text;not null;default:''"`
	Secret      string   `gorm:"type:text;not null"`
	Partition   string   `gorm:"type:text;not null"`
	Default     bool     `gorm:"type:boolean;not null;default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewAuthClient(data ctype.Dict) *AuthClient {
	return &AuthClient{
		Uid:         data["Uid"].(string),
		Description: data["Description"].(string),
		Secret:      data["Secret"].(string),
		Partition:   data["Partition"].(string),
		Default:     data["Default"].(bool),
	}
}

type Tenant struct {
	ID           uint
	AuthClientID uint
	AuthClient   *AuthClient
	Uid          string `gorm:"type:text;not null;unique"`
	Title        string `gorm:"type:text;not null"`
	Avatar       string `gorm:"type:text;not null;default:''"`
	AvatarStr    string `gorm:"type:text;not null;default:''"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewTenant(data ctype.Dict) *Tenant {
	return &Tenant{
		AuthClientID: data["AuthClientID"].(uint),
		Uid:          data["Uid"].(string),
		Title:        data["Title"].(string),
		Avatar:       data["Avatar"].(string),
		AvatarStr:    data["AvatarStr"].(string),
	}
}

type User struct {
	ID          uint `json:"id"`
	TenantID    uint `gorm:"not null;uniqueIndex:idx_users_tenant_uid;uniqueIndex:idx_users_tenant_email" json:"tenant_id"`
	Tenant      *Tenant
	TenantTmpID *uint          `json:"tenant_tmp_id"`
	Roles       []Role         `gorm:"many2many:users_roles;" json:"roles"`
	Uid         string         `gorm:"type:text;not null;uniqueIndex:idx_users_tenant_uid" json:"uid"`
	Email       string         `gorm:"type:text;not null;uniqueIndex:idx_users_tenant_email" json:"email"`
	Mobile      *string        `gorm:"type:text" json:"mobile"`
	FirstName   string         `gorm:"type:text;not null;default:''" json:"first_name"`
	LastName    string         `gorm:"type:text;not null;default:''" json:"last_name"`
	Avatar      string         `gorm:"type:text;not null;default:''" json:"avatar"`
	AvatarStr   string         `gorm:"type:text;not null;default:''" json:"avatar_str"`
	ExtraInfo   datatypes.JSON `gorm:"type:json;not null;default:'{}'" json:"extra_info"`
	Admin       bool           `gorm:"type:boolean;not null;default:false" json:"admin"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

func NewUser(data ctype.Dict) *User {
	extraInfoJSON, err := json.Marshal(data["ExtraInfo"])
	if err != nil {
		panic("Failed to marshal ExtraInfo")
	}
	return &User{
		TenantID:    data["TenantID"].(uint),
		TenantTmpID: data["TenantTmpID"].(*uint),
		Uid:         data["Uid"].(string),
		Email:       data["Email"].(string),
		Mobile:      data["Mobile"].(*string),
		FirstName:   data["FirstName"].(string),
		LastName:    data["LastName"].(string),
		Avatar:      data["Avatar"].(string),
		AvatarStr:   data["AvatarStr"].(string),
		ExtraInfo:   datatypes.JSON(extraInfoJSON),
		Admin:       data["Admin"].(bool),
	}
}

type Role struct {
	ID        uint
	Users     []User `gorm:"many2many:users_roles;"`
	Pems      []Pem  `gorm:"many2many:roles_pems;"`
	TenantID  uint   `gorm:"not null;uniqueIndex:idx_roles_tenant_title"`
	Tenant    *Tenant
	Title     string `gorm:"type:text;not null;uniqueIndex:idx_roles_tenant_title"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewRole(data ctype.Dict) *Role {
	return &Role{
		TenantID: data["TenantID"].(uint),
		Title:    data["Title"].(string),
	}
}

type Pem struct {
	ID     uint
	Roles  []Role `gorm:"many2many:roles_pems;"`
	Title  string `gorm:"type:text;not null"`
	Module string `gorm:"type:text;not null;uniqueIndex:idx_pems_module_action"`
	Action string `gorm:"type:text;not null;uniqueIndex:idx_pems_module_action"`
}

func NewPem(data ctype.Dict) *Pem {
	return &Pem{
		Title:  data["Title"].(string),
		Module: data["Module"].(string),
		Action: data["Action"].(string),
	}
}
