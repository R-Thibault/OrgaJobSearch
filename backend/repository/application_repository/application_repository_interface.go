package applicationrepository

import "github.com/R-Thibault/OrgaJobSearch/backend/models"

type ApplicationRepositoryInterface interface {
	SaveApplication(appDatas models.Application) error
}
