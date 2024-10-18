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

func (s *OtpGeneratorService) GenerateOTP(user *models.User) (otp models.OTP) {
	// Set the seed for the random number generator
	rand.Seed(uint64(time.Now().UnixNano()))
	OtpRandCode := rand.Intn(900000) + 100000
	OtpRandCodeToString := strconv.Itoa(OtpRandCode)
	otpGenerated := models.OTP{
		UserID:        user.ID,
		OtpCode:       OtpRandCodeToString,
		OtpExpiration: time.Now().Add(60 * time.Minute), // 1heure
		OtpType:       "emailValidation",
		OtpAttempts:   0,
		IsValid:       true,
	}

	return otpGenerated
}
