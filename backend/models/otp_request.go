package models

type SendOTPRequest struct {
	Email   string `json:"email" binding:"required"`
	OtpCode string `json:"otpCode" binding:"required"`
}
