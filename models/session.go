package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Identificator string `gorm:"index"`
	UserID        uint
	User          *User `gorm:"foreignKey:UserID"`
	FirstIP       string
	LastIP        string
	ExpiresAt     time.Time `gorm:"notnull;default:CURRENT_TIMESTAMP"`
}
