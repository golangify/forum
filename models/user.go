package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"unique;notnull"`
	PasswordHash string
	Roles        datatypes.JSONSlice[string]
}
