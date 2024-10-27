package schema

import (
	"src/common/ctype"
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
	ID          uint      `json:"id" gorm:"primary_key"`
	Key         string    `gorm:"type:text;not null;unique" json:"key"`
	Value       string    `gorm:"type:text;not null;default:''" json:"value"`
	Description string    `gorm:"type:text;not null;default:''" json:"description"`
	DataType    string    `gorm:"type:text;not null;default:'STRING';check:data_type IN ('STRING', 'INTEGER', 'FLOAT', 'BOOLEAN', 'DATE', 'DATETIME')" json:"data_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewVariable(data ctype.Dict) *Variable {
	return &Variable{
		Key:         data["Key"].(string),
		Value:       data["Value"].(string),
		Description: data["Description"].(string),
		DataType:    data["DataType"].(string),
	}
}
