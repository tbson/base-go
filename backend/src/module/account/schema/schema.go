package schema

import (
	"encoding/json"
	"time"

	"src/common/ctype"
	"src/util/dictutil"

	"gorm.io/datatypes"
)

type AuthClient struct {
	ID          uint     `gorm:"primaryKey" json:"id"`
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
	ID           uint `gorm:"primaryKey" json:"id"`
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
	ID          uint `gorm:"primaryKey" json:"id"`
	TenantID    uint `gorm:"not null;uniqueIndex:idx_users_tenant_external;uniqueIndex:idx_users_tenant_email" json:"tenant_id"`
	Tenant      *Tenant
	TenantTmpID *uint          `json:"tenant_tmp_id"`
	Sub         *string        `gorm:"type:text;default:null;unique" json:"sub"`
	Roles       []Role         `gorm:"many2many:users_roles;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"roles"`
	ExternalID  string         `gorm:"type:text;not null;uniqueIndex:idx_users_tenant_external" json:"uid"`
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
		TenantID:    dictutil.GetValue[uint](data, "TenantID"),
		TenantTmpID: dictutil.GetValue[*uint](data, "TenantTmpID"),
		ExternalID:  dictutil.GetValue[string](data, "ExternalID"),
		Email:       dictutil.GetValue[string](data, "Email"),
		Mobile:      dictutil.GetValue[*string](data, "Mobile"),
		FirstName:   dictutil.GetValue[string](data, "FirstName"),
		LastName:    dictutil.GetValue[string](data, "LastName"),
		Avatar:      dictutil.GetValue[string](data, "Avatar"),
		AvatarStr:   dictutil.GetValue[string](data, "AvatarStr"),
		ExtraInfo:   datatypes.JSON(extraInfoJSON),
		Admin:       dictutil.GetValue[bool](data, "Admin"),
	}
}

type Role struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Users     []User `gorm:"many2many:users_roles;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	Pems      []Pem  `gorm:"many2many:roles_pems;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
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
	ID     uint   `gorm:"primaryKey" json:"id"`
	Roles  []Role `gorm:"many2many:roles_pems;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
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
