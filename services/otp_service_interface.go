package services

type OTPServiceInterface interface {
	GenerateOTP(email string) (otpCode string, err error)
}
