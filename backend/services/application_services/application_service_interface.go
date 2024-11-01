package applicationservices

import "github.com/R-Thibault/OrgaJobSearch/backend/models"

type ApplicationServiceInterface interface {
	SaveApplication(userID uint, appData models.Application) error
	UpdateApplication(userID uint, appData models.Application) (*models.Application, error)
	GetApplicationByID(userID uint, applicationID uint) (*models.Application, error)
	GetApplicationsByUserID(userID uint) ([]*models.Application, error)
	DeleteApplication(userID uint, applicationID uint) error
}
