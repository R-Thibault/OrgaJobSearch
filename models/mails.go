package models

type OTPMail struct {
	ToEmail          string
	Subject          string
	plainTextContent string
	htmlContent      string
	OTPCode          string
}
