package services

type OTPServiceInterface interface {
	GenerateOTP(email string) (otpCode string, err error)
	VerifyOTP(email string, otpCode string) error
}
