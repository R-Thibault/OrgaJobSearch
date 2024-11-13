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

// UpdateOTP updates an existing OTP record in the database.
// It takes a pointer to a models.OTP struct as input and returns the updated OTP record or an error.
//
// Parameters:
//   - newOTP: A pointer to the models.OTP struct containing the OTP details to be updated.
//
// Returns:
//   - *models.OTP: A pointer to the updated OTP record.
//   - error: An error if the update operation fails.
//
// Errors:
//   - Returns an error if the database connection is nil.
//   - Returns an error if the OtpCode is "0".
//   - Returns an error if the OtpType is empty.
//   - Returns gorm.ErrRecordNotFound if the OTP record is not found in the database.
//   - Returns any other error encountered during the update operation.
func (r *OTPRepository) UpdateOTP(newOTP *models.OTP) (*models.OTP, error) {
	if r.db == nil {
		return &models.OTP{}, errors.New("database connection is nil")
	}

	if newOTP.OtpCode == "0" {
		return &models.OTP{}, errors.New("OtpCode is empty")
	}
	if newOTP.OtpType == "" {
		return &models.OTP{}, errors.New("otpType is empty")
	}

	result := r.db.Clauses(clause.Returning{}).Save(newOTP)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &models.OTP{}, gorm.ErrRecordNotFound
		}
		return &models.OTP{}, result.Error
	}
	return newOTP, nil
}
