package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/R-Thibault/OrgaJobSearch/models"
	"github.com/R-Thibault/OrgaJobSearch/services"
	"github.com/gin-gonic/gin"
)

type OTPController struct {
	OTPService    *services.OTPService
	MailerService *services.MailerService
}

func NewOTPController(OTPService *services.OTPService, MailerService *services.MailerService) *OTPController {

	return &OTPController{OTPService: OTPService, MailerService: MailerService}
}

func (u *OTPController) GenerateOTP(c *gin.Context) {
	var creds models.Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		// If the input is invalid, respond with an error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	Otp, err := u.OTPService.GenerateOTP(creds.Email)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	log.Printf("%v\n", Otp)
	// Logique envoie mail ici

	c.JSON(http.StatusOK, gin.H{"otp": "OTP generated successfully"})

	go func() {
		// Create the request payload
		requestBody := models.SendOTPRequest{
			Email:   creds.Email,
			OtpCode: Otp,
		}

		// Convert the payload to JSON
		jsonData, err := json.Marshal(requestBody)
		if err != nil {
			log.Printf("Error marshalling JSON: %v", err)
			return
		}
		client := &http.Client{}
		req, err := http.NewRequest("POST", "http://localhost:8080/send-otp", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("Error creating request to send OTP: %v", err)
			return
		}
		req.Header.Set("Content-type", "application/json")
		resp, err := client.Do(req)
		if err != nil || resp.StatusCode != http.StatusOK {
			log.Printf("Error sending OTP: %v", err)
		} else {
			log.Printf("OTP send successfully")
		}
	}()
}

func (u *OTPController) SendOTP(c *gin.Context) {
	var request models.SendOTPRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		// If the input is invalid, respond with an error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	err := u.MailerService.SendOTPMail(request.Email, request.OtpCode)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP Successfully send"})
}

func (ctrl *OTPController) ValidateOTP(c *gin.Context) {
	var request struct {
		OTP int `json:"otp"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Here you would typically check the OTP against a stored value
	// For simplicity, we'll assume any OTP is valid
	c.JSON(http.StatusOK, gin.H{"message": "OTP is valid"})
}

func RegisterRoutes(router *gin.Engine) {
	otpController := NewOTPController()
	router.POST("/generate-otp", otpController.GenerateOTP)
	router.POST("/validate-otp", otpController.ValidateOTP)
}