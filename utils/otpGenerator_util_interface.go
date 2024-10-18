package utils

import "github.com/R-Thibault/OrgaJobSearch/models"

type OtpGeneratorServiceInterface interface {
	GenerateOTP(user *models.User) models.OTP
}
