package models

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	SectionID uint
	Section   *Section `gorm:"foreignKey:SectionID"`
	UserID    uint
	User      *User `gorm:"foreignKey:UserID"`
	Title     string
	Body      string
}
