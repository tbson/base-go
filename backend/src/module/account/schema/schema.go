package schema

import (
	"time"

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

type Tenant struct {
	ID           uint
	AuthClientID *uint
	AuthClient   *AuthClient
	Uid          string `gorm:"type:text;not null;unique"`
	Title        string `gorm:"type:text;not null"`
	Avatar       string `gorm:"type:text;not null;default:''"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type User struct {
	ID          uint
	TenantID    uint `gorm:"not null;uniqueIndex:idx_users_tenant_uid;uniqueIndex:idx_users_tenant_email"`
	Tenant      Tenant
	TenantTmpID *uint
	Uid         string         `gorm:"type:text;not null;uniqueIndex:idx_users_tenant_uid"`
	Email       string         `gorm:"type:text;not null;uniqueIndex:idx_users_tenant_email"`
	Mobile      *string        `gorm:"type:text"`
	FirstName   string         `gorm:"type:text;not null;default:''"`
	LastName    string         `gorm:"type:text;not null;default:''"`
	Avatar      string         `gorm:"type:text;not null;default:''"`
	AvatarStr   string         `gorm:"type:text;not null;default:''"`
	ExtraInfo   datatypes.JSON `gorm:"type:json;not null;default:'{}'"`
	Admin       bool           `gorm:"type:boolean;not null;default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
