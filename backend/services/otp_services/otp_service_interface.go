package services

import (
	"time"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
)

type OTPServiceInterface interface {
	GenerateOTP(userID uint, otpType string, expirationTime time.Time) (otpCode string, err error)
	VerifyOTPGiven(email string, otpType string, otpCode string) error
	VerifyOTPCode(otpCode string, otpType string) (*models.OTP, error)
	CheckAndRefreshOTPCode(userID uint, otpType string, expirationTime time.Time) (otpCode string, err error)
}
