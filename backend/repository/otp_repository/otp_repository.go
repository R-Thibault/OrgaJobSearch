package repository

import (
	"errors"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (r *OTPRepository) GetOTPCodeByUserIDandType(userID uint, otpType string) (*models.OTP, error) {
	if r.db == nil {
		return &models.OTP{}, errors.New("database connection is nil")
	}

	if userID == 0 {
		return &models.OTP{}, errors.New("userID is zero")
	}

	var otp models.OTP
	result := r.db.Where("user_id = ? AND otp_type = ?", userID, otpType).First(&otp)
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

func (r *OTPRepository) GetOTPByCode(otpCode string, otpType string) (*models.OTP, error) {
	if r.db == nil {
		return &models.OTP{}, errors.New("database connection is nil")
	}
	if otpCode == "0" {
		return &models.OTP{}, errors.New("OtpCode is empty")
	}
	var otp models.OTP
	result := r.db.Where("otp_code = ? AND otp_type = ?", otpCode, otpType).First(&otp)
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

func (r *OTPRepository) UpdateOTPCode(otpID uint, otpCode string, otpType string) (*models.OTP, error) {
	if r.db == nil {
		return &models.OTP{}, errors.New("database connection is nil")
	}
	if otpID == 0 {
		return &models.OTP{}, errors.New("otpID is empty")
	}
	if otpCode == "0" {
		return &models.OTP{}, errors.New("OtpCode is empty")
	}
	if otpType == "" {
		return &models.OTP{}, errors.New("otpType is empty")
	}
	updatedOTP := &models.OTP{
		Model: gorm.Model{
			ID: otpID,
		},
		OtpCode: otpCode,
	}
	result := r.db.Clauses(clause.Returning{}).Save(updatedOTP)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &models.OTP{}, gorm.ErrRecordNotFound
		}
		return &models.OTP{}, result.Error
	}
	return updatedOTP, nil
}
