package repository

import (
	"errors"

	"github.com/R-Thibault/OrgaJobSearch/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) SaveUser(user models.User) error {
	//Save user in DB
	return r.db.Create(&user).Error
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	if r.db == nil {
		return nil, errors.New("database connection is nil")
	}

	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	var user models.User
	result := r.db.Debug().Where("email = ?", email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}
