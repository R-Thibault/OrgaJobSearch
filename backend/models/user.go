package models

import "gorm.io/gorm"

// User table for Database
type User struct {
	gorm.Model
	FirstName      string `gorm:"size:255"`
	LastName       string `gorm:"size:255"`
	Email          string `gorm:"size:255;unique;index;not null"`
	HashedPassword string `gorm:"size:255;"`
	UserStatus     string `gorm:"size:255; not null"`
	UserUUID       string `gorm:"size:36;index;unique;not null"`
	EmailIsValide  bool   `gorm:"default:false"`
	Otps           []OTP
	Applications   []Application `gorm:"foreignkey:UserID"`
}

type UserProfileUpdate struct {
	UserFirstName string
	UserLastName  string
	Email         string
}
