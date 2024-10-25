package models

import "gorm.io/gorm"

// User table for Database
type User struct {
	gorm.Model
	Name            string `gorm:"size:255"`
	Email           string `gorm:"size:255;unique;index;not null"`
	HashedPassword  string `gorm:"size:255;"`
	UserStatus      string `gorm:"size:255"`
	UserUUID        string `gorm:"size:36"`
	EmailIsValide   bool   `gorm:"default:false"`
	Otps            []OTP
	Roles           []Role `gorm:"many2many:user_role;"`
	CareerSupportID *uint  // Foreign key for self-referencing the Career Support
	CareerSupport   *User  `gorm:"foreignkey:CareerSupportID"` // Reference to the Career Support (optional)
	Jobseekers      []User `gorm:"foreignkey:CareerSupportID"` // List of Jobseekers being followed
}
