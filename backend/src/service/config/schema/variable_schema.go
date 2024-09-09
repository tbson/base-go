package schema

import (
	"gorm.io/gorm"
)

type Variable struct {
	gorm.Model
	Key         string `gorm:"type:varchar(255);not null;unique"`
	Value       string `gorm:"type:varchar(255);not null;default:''"`
	Description string `gorm:"type:varchar(255);not null;default:''"`
	Type        uint8  `gorm:"not null;default:1"`
}
