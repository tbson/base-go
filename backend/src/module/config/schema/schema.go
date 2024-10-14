package schema

import (
	"src/util/iterutil"
	"time"
)

var TypeDict = iterutil.FieldEnum{
	"STRING",
	"INTEGER",
	"FLOAT",
	"BOOLEAN",
	"DATE",
	"DATETIME",
}

var TypeOptions = iterutil.GetFieldOptions(TypeDict)

type Variable struct {
	ID          uint
	Key         string `gorm:"type:text;not null;unique"`
	Value       string `gorm:"type:text;not null;default:''"`
	Description string `gorm:"type:text;not null;default:''"`
	DataType    string `gorm:"type:text;not null;default:'STRING';check:data_type IN ('STRING', 'INTEGER', 'FLOAT', 'BOOLEAN', 'DATE', 'DATETIME')"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
