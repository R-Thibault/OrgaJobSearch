package models

import "gorm.io/gorm"

// Application table for database

type Application struct {
	gorm.Model
	UserID      uint
	Url         string `gorm:"not null"`
	Title       string `gorm:"size:255; not null"`
	Company     string `gorm:"size:255"`
	Location    string `gorm:"size:255"`
	Description string
	Salary      string `gorm:"size:255"`
	JobType     string `gorm:"size:255"`
	Applied     bool   `gorm:"default:true"`
	Response    bool   `gorm:"default:false"`
	FolloUp     bool   `gorm:"default:false"`
}
