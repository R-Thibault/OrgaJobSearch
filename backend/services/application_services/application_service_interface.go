package applicationservices

import "github.com/R-Thibault/OrgaJobSearch/backend/models"

type ApplicationServiceInterface interface {
	SaveApplication(userID uint, appData models.Application) error
}
