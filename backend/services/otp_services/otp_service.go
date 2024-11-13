package services

import (
	"errors"
	"time"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	otpRepository "github.com/R-Thibault/OrgaJobSearch/backend/repository/otp_repository"
	userRepository "github.com/R-Thibault/OrgaJobSearch/backend/repository/user_repository"
	otpGeneratorUtils "github.com/R-Thibault/OrgaJobSearch/backend/utils/otpGenerator_util"
)

type OTPService struct {
	userRepo userRepository.UserRepositoryInterface
	OTPRepo  otpRepository.OTPRepositoryInterface
	OTPUtil  otpGeneratorUtils.OtpGeneratorServiceInterface
}

func NewOTPService(userRepo userRepository.UserRepositoryInterface, OTPRepo otpRepository.OTPRepositoryInterface, OTPUtil otpGeneratorUtils.OtpGeneratorServiceInterface) *OTPService {
	return &OTPService{userRepo: userRepo, OTPRepo: OTPRepo, OTPUtil: OTPUtil}
}

var _ OTPServiceInterface = &OTPService{}

func (s *OTPService) GenerateOTP(userID uint, otpType string, expirationTime time.Time) (otpCode string, err error) {

	user, err := s.userRepo.GetUserByID(userID)
	if err != nil || user == nil {
		return "", errors.New("user not found")
	}

	otp := s.OTPUtil.GenerateOTP(user, otpType, expirationTime)

	otpCodeGenerated, err := s.OTPRepo.SaveOTP(otp)
	if err != nil {
		return "", errors.New("Problem during OTP creation")
	}

	return otpCodeGenerated, nil
}

func (s *OTPService) VerifyOTPGiven(email string, otpType string, otpCode string) error {
	// fetch user by email
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil || user == nil {
		return errors.New("user not found")
	}
	// Fetch OTP associated with the user
	otpSaved, err := s.OTPRepo.GetOTPCodeByUserIDandType(user.ID, otpType)
	if err != nil || otpSaved == nil {
		return errors.New("OTP not found")
	}

	// Verify OTP
	if otpSaved.OtpCode == otpCode {
		if otpSaved.OtpExpiration.After(time.Now()) {
			return nil
		} else {
			return errors.New("OTP expired")
		}
	} else {
		return errors.New("OTP codes do not match")
	}

}

func (s *OTPService) CheckAndRefreshOTPCode(userID uint, otpType string, expirationTime time.Time) (otpCode string, err error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil || user == nil {
		return "", errors.New("user not found")
	}
	otp, err := s.OTPRepo.GetOTPCodeByUserIDandType(user.ID, otpType)
	if err != nil || otp == nil {
		return "", errors.New("OTP doesn't exists")
	}

	if time.Now().After(otp.OtpExpiration) {
		newOTPCode := s.OTPUtil.GenerateOTP(user, otpType, expirationTime)
		newOtp, err := s.OTPRepo.UpdateOTP(&newOTPCode)
		if err != nil {
			return "", errors.New("Error during OTP update")
		}
		return newOtp.OtpCode, nil
	}
	return otp.OtpCode, nil
}

func (s *OTPService) VerifyOTPCode(otpCode string, otpType string) (*models.OTP, error) {
	if otpCode == "" {
		return &models.OTP{}, errors.New("Otp code can't be emtpy")
	}
	hours2 := 2 * time.Hour
	otpSaved, err := s.OTPRepo.GetOTPByCode(otpCode, otpType)
	if err != nil || otpSaved == nil {
		return &models.OTP{}, errors.New("OTP not found")
	} else if time.Since(otpSaved.UpdatedAt) >= hours2 {
		return &models.OTP{}, errors.New("OTP not valid")
	}
	return otpSaved, nil
}
