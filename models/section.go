package models

import (
	"html/template"

	"gorm.io/gorm"
)

type Section struct {
	gorm.Model
	UserID        uint
	User          *User `gorm:"foreignKey:UserID"`
	Title         string
	Body          string
	BodyHTML      template.HTML
	TopicsCount   uint
	CommentsCount uint
}
