package applicationservices

import (
	"errors"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	applicationrepository "github.com/R-Thibault/OrgaJobSearch/backend/repository/application_repository"
)

type ApplicationService struct {
	ApplicationRepo applicationrepository.ApplicationRepositoryInterface
}

func NewApplicationService(ApplicationRepo applicationrepository.ApplicationRepositoryInterface) *ApplicationService {
	return &ApplicationService{ApplicationRepo: ApplicationRepo}
}

func (s *ApplicationService) SaveApplication(userID uint, appData models.Application) error {
	if appData.Title == "" {
		return errors.New("Title can't be empty")
	}
	appData.UserID = userID

	return s.ApplicationRepo.SaveApplication(appData)
}
