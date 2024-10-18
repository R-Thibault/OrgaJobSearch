package services

import (
	"errors"
	"time"

	otpRepository "github.com/R-Thibault/OrgaJobSearch/repository/otp_repository"
	userRepository "github.com/R-Thibault/OrgaJobSearch/repository/user_repository"
	otpGeneratorUtils "github.com/R-Thibault/OrgaJobSearch/utils/otpGenerator_util"
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

func (s *OTPService) GenerateOTP(email string) (otpCode string, err error) {

	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil || user == nil {
		return "", errors.New("user not found")
	}

	otp := s.OTPUtil.GenerateOTP(user)

	otpCodeGenerated, err := s.OTPRepo.SaveOTP(otp)
	if err != nil {
		return "", errors.New("Problem during OTP creation")
	}

	return otpCodeGenerated, nil
}

func (s *OTPService) VerifyOTP(email string, otpCode string) error {
	// fetch user by email
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil || user == nil {
		return errors.New("user not found")
	}

	// Fetch OTP associated with the user
	otpSaved, err := s.OTPRepo.GetOTPCodeByUserID(user.ID)
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
