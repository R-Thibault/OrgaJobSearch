package models

import "gorm.io/gorm"

// User table for Database
type User struct {
	gorm.Model
	Name           string `gorm:"size:255"`
	Email          string `gorm:"size:255;unique;index;not null"`
	HashedPassword string `gorm:"size:255;"`
	EmailIsValide  bool   `gorm:"default:false"`
	Otps           []OTP
	Roles          []Role `gorm:"many2many:user_role;"`
}
