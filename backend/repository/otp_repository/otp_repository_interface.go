package repository

import "github.com/R-Thibault/OrgaJobSearch/backend/models"

// UserRepositoryInterface defines the methods for interacting with users in the database.
type OTPRepositoryInterface interface {
	SaveOTP(otp models.OTP) (otpCode string, err error)
	GetOTPCodeByUserIDandType(userID uint, otpType string) (*models.OTP, error)
	GetOTPByCode(otpCode string, otpType string) (*models.OTP, error)
	UpdateOTPCode(otpID uint, otpCode string, otpType string) (*models.OTP, error)
}
