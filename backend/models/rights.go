package models

import "gorm.io/gorm"

// Right table for database
type Role struct {
	gorm.Model
	RoleName    string `gorm:"size:255;unique; not null"`
	Description string `gorm:"nut null"`
}
