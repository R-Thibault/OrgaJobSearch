package repository

import (
	"errors"

	"github.com/R-Thibault/OrgaJobSearch/models"
	"gorm.io/gorm"
)

type OTPRepository struct {
	db *gorm.DB
}

func NewOTPRepository(db *gorm.DB) *OTPRepository {
	return &OTPRepository{db: db}
}

var _ OTPRepositoryInterface = &OTPRepository{}

func (r *OTPRepository) SaveOTP(otp models.OTP) (otpCode string, err error) {
	err = r.db.Create(&otp).Error
	if err != nil {
		return "", err
	}
	return otp.OtpCode, nil
}

func (r *OTPRepository) GetOTPCodeByUserID(userID uint) (*models.OTP, error) {
	if r.db == nil {
		return &models.OTP{}, errors.New("database connection is nil")
	}

	if userID == 0 {
		return &models.OTP{}, errors.New("userID is zero")
	}

	var otp models.OTP
	result := r.db.Where("user_id = ? AND otp_type = ?", userID, "emailValidation").First(&otp)
	if result.Error != nil {
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil, gorm.ErrRecordNotFound
			}
			return nil, result.Error
		}
	}

	return &otp, nil
}
