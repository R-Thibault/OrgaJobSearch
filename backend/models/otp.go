package models

import (
	"time"

	"gorm.io/gorm"
)

type OTP struct {
	gorm.Model
	UserID        uint
	OtpCode       string `gorm:"size:10;not null"`
	OtpExpiration time.Time
	OtpType       string `gorm:"size:255"`
	OtpAttempts   int
	IsValid       bool `gorm:"default:true"` //To know if Otp is still valid
}
