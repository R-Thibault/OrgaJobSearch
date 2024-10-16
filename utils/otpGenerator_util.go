package utils

import (
	"time"

	"github.com/R-Thibault/OrgaJobSearch/models"
)

type OtpGeneratorService struct {
}

// NewOtpGeneratorService creates a new instance of OtpGeneratorService.
func NewOtpGeneratorService() *OtpGeneratorService {
	return &OtpGeneratorService{}
}

var _ OtpGeneratorServiceInterface = &OtpGeneratorService{}

func (s *OtpGeneratorService) GenerateOTP(user *models.User) (otp models.OTP) {
	otpGenerated := models.OTP{
		UserID:        user.ID,
		OtpCode:       "123456",
		OtpExpiration: time.Now().Add(60 * time.Minute),
		OtpType:       "login",
		OtpAttempts:   0,
		IsValid:       true,
	}

	return otpGenerated
}
