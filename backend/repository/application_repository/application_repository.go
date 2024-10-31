package applicationrepository

import (
	"errors"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	"gorm.io/gorm"
)

type ApplicationRepository struct {
	db *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) *ApplicationRepository {
	return &ApplicationRepository{db: db}
}

func (r *ApplicationRepository) SaveApplication(appDatas models.Application) error {
	if r.db == nil {
		return errors.New("database connection is nil")
	}
	return r.db.Create(&appDatas).Error

}
