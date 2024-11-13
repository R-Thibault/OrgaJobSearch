package utils

import (
	"strconv"
	"time"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	"golang.org/x/exp/rand"
)

type OtpGeneratorService struct {
}

// NewOtpGeneratorService creates a new instance of OtpGeneratorService.
func NewOtpGeneratorService() *OtpGeneratorService {
	return &OtpGeneratorService{}
}

var _ OtpGeneratorServiceInterface = &OtpGeneratorService{}

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
//     number of attempts, and validity status.
func (s *OtpGeneratorService) GenerateOTP(user *models.User, OtpType string, expirationTime time.Time) (otp models.OTP) {
	// Set the seed for the random number generator
	rand.Seed(uint64(time.Now().UnixNano()))
	OtpRandCode := rand.Intn(900000) + 100000
	OtpRandCodeToString := strconv.Itoa(OtpRandCode)
	otpGenerated := models.OTP{
		UserID:        user.ID,
		OtpCode:       OtpRandCodeToString,
		OtpExpiration: expirationTime, // 1heure
		OtpType:       OtpType,
		OtpAttempts:   0,
		IsValid:       true,
	}

	return otpGenerated
}
