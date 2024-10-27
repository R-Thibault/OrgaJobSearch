package repository

import "github.com/R-Thibault/OrgaJobSearch/backend/models"

// UserRepositoryInterface defines the methods for interacting with users in the database.
type OTPRepositoryInterface interface {
	SaveOTP(otp models.OTP) (otpCode string, err error)
	GetOTPCodeByUserID(userID uint) (*models.OTP, error)
	GetOTPByCode(otpCode string, otpType string) (*models.OTP, error)
}
