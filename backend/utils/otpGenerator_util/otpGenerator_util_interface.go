package utils

import "github.com/R-Thibault/OrgaJobSearch/backend/models"

type OtpGeneratorServiceInterface interface {
	GenerateOTP(user *models.User) models.OTP
}
