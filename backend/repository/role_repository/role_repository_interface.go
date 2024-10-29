package rolerepository

import "github.com/R-Thibault/OrgaJobSearch/backend/models"

type RoleRepositoryInterface interface {
	GetRoleByName(roleName string) (*models.Role, error)
}
