package applicationrepository

import "github.com/R-Thibault/OrgaJobSearch/backend/models"

type ApplicationRepositoryInterface interface {
	SaveApplication(appDatas models.Application) error
	UpdateApplication(appDatas models.Application) (*models.Application, error)
	GetApplicationByID(applicationID uint) (*models.Application, error)
	GetApplicationsByUserID(userID uint) ([]*models.Application, error)
	DeleteApplication(application models.Application) error
}
