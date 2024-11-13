package utils

import (
	"time"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
)

type OtpGeneratorServiceInterface interface {
	// GenerateOTP generates a one-time password (OTP) for a given user.
	// It sets the seed for the random number generator, generates a 6-digit OTP,
	// and returns an OTP model with the provided expiration time and OTP type.
	//
	// Parameters:
	//   - user: A pointer to the User model for whom the OTP is being generated.
	//   - OtpType: A string representing the type of OTP being generated.
	//   - expirationTime: A time.Time value representing the expiration time of the OTP.
	//
	// Returns:
	//   - otp: An OTP model containing the generated OTP code, user ID, expiration time, OTP type,
	//          number of attempts, and validity status.
	GenerateOTP(user *models.User, OtpType string, expirationTime time.Time) (otp models.OTP)
}
