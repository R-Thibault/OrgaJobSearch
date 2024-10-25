package invitationservices

import (
	"errors"
	"fmt"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	userRepository "github.com/R-Thibault/OrgaJobSearch/backend/repository/user_repository"
)

type Invitationservice struct {
	UserRepo userRepository.UserRepositoryInterface
}

func NewInvitationService(UserRepo userRepository.UserRepositoryInterface) *Invitationservice {
	return &Invitationservice{UserRepo: UserRepo}
}

func (s *Invitationservice) VerifyPersonnalInvitationTokenData(token models.JWTToken) (email string, err error) {
	if *token.Body == "" {
		return "", errors.New("empty UUID")
	}
	existingUser, err := s.UserRepo.GetUserByUUID(*token.Body)

	if err != nil {
		return "", fmt.Errorf("failed to check if user exists: %w", err)
	}
	return existingUser.Email, nil

}
