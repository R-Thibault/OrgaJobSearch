package invitationservices

import (
	"errors"
	"fmt"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	userRepository "github.com/R-Thibault/OrgaJobSearch/backend/repository/user_repository"
	otpServices "github.com/R-Thibault/OrgaJobSearch/backend/services/otp_services"
)

type InvitationService struct {
	UserRepo   userRepository.UserRepositoryInterface
	OTPService otpServices.OTPServiceInterface
}

func NewInvitationService(UserRepo userRepository.UserRepositoryInterface, OTPService otpServices.OTPServiceInterface) *InvitationService {
	return &InvitationService{UserRepo: UserRepo, OTPService: OTPService}
}

func (s *InvitationService) VerifyPersonnalInvitationTokenData(token models.JWTToken) (email string, err error) {
	if *token.Body == "" {
		return "", errors.New("empty body")
	}
	existingUser, err := s.UserRepo.GetUserByUUID(*token.Body)

	if err != nil {
		return "", fmt.Errorf("failed to check if user exists: %w", err)
	}
	return existingUser.Email, nil
}

func (s *InvitationService) VerifyGlobalInvitationTokenData(token models.JWTToken) error {
	if *token.Body == "" {
		return errors.New("empty body")
	}
	err := s.OTPService.VerifyOTPForGlobalInvitation(*token.Body, *token.TokenType)
	if err != nil {
		return errors.New("Otp invalid")
	}
	return nil
}
