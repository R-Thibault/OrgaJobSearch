package services

type OTPServiceInterface interface {
	GenerateOTP(email string, otpType string) (otpCode string, err error)
	VerifyOTPForGlobalInvitation(otpCode string, otpType string) error
}
