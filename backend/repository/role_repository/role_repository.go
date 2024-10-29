package rolerepository

import (
	"errors"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) GetRoleByName(roleName string) (*models.Role, error) {
	if roleName == "" {
		return &models.Role{}, errors.New("Role name can't be empty")
	}
	var role models.Role
	result := r.db.Unscoped().Where("role_name = ?", roleName).First(&role)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	return &role, nil
}
