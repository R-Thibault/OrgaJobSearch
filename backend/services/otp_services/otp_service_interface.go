package services

type OTPServiceInterface interface {
	GenerateOTP(userID uint, otpType string) (otpCode string, err error)
	VerifyOTP(email string, otpCode string) error
	VerifyOTPForGlobalInvitation(otpCode string, otpType string) error
	CheckOTPCodeForGlobalInvitation(userID uint, otpType string) (otpCode string, err error)
}
