package repository

import "github.com/R-Thibault/OrgaJobSearch/backend/models"

// UserRepositoryInterface defines the methods for interacting with users in the database.
type OTPRepositoryInterface interface {
	SaveOTP(otp models.OTP) (otpCode string, err error)
	GetOTPCodeByUserIDandType(userID uint, otpType string) (*models.OTP, error)
	GetOTPByCode(otpCode string, otpType string) (*models.OTP, error)
	// UpdateOTP updates an existing OTP entry with new values.
	// Parameters:
	//   - newOTP: An OTP struct containing updated values.
	//       - newOTP.OtpCode must not be empty.
	//       - newOTP.OtpType must not be empty.
	//       - newOTP.OtpExpiration should be set to the desired expiration time.
	// Returns:
	//   - A pointer to the updated OTP struct if the operation is successful.
	//   - An error if the database connection is nil, if otpID or required fields in newOTP are empty,
	//     or if any issue occurs during the database operation.
	UpdateOTP(newOTP *models.OTP) (*models.OTP, error)
}
