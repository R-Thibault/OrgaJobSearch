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

func (s *ApplicationService) UpdateApplication(userID uint, appData models.Application) (*models.Application, error) {
	if appData.UserID != userID {
		return &models.Application{}, errors.New("The application doesn't belong to this user")
	}
	return s.ApplicationRepo.UpdateApplication(appData)
}

func (s *ApplicationService) GetApplicationByID(userID uint, applicationID uint) (*models.Application, error) {
	if userID == 0 || applicationID == 0 {
		return &models.Application{}, errors.New("userID or ApplicationID can't be null")
	}
	application, err := s.ApplicationRepo.GetApplicationByID(applicationID)
	if err != nil || application == nil {
		return &models.Application{}, errors.New("Can't find application")
	}
	if application.UserID != userID {
		return &models.Application{}, errors.New("Application not created by user")
	}
	return application, nil
}

func (s *ApplicationService) GetApplicationsByUserID(userID uint) ([]*models.Application, error) {
	if userID == 0 {
		return nil, errors.New("userID or ApplicationID can't be null")
	}
	applications, err := s.ApplicationRepo.GetApplicationsByUserID(userID)
	if err != nil {
		return nil, errors.New("Can't find applications with this userID")
	}
	return applications, nil
}

func (s *ApplicationService) DeleteApplication(userID uint, applicationID uint) error {
	if userID == 0 || applicationID == 0 {
		return errors.New("UserID or ApplicationID can't be null")
	}
	application, err := s.ApplicationRepo.GetApplicationByID(applicationID)
	if err != nil || application == nil {
		return errors.New("Can't find application")
	}
	if application.UserID != userID {
		return errors.New("Application not created by user")
	}
	deleteErr := s.ApplicationRepo.DeleteApplication(*application)
	if deleteErr != nil {
		return errors.New("Error during supression process")
	}
	return nil
}
