package schema

type Variable struct {
	ID          uint64 `gorm:"primary_key"`
	Value       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:varchar(255);not null;default:''"`
	Type        uint8  `gorm:"not null;default:1"`
}
