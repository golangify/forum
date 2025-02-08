package models

import "gorm.io/gorm"

type Section struct {
	gorm.Model
	UserID uint
	User   *User `gorm:"foreignKey:UserID"`
	Title  string
	Body   string
}
