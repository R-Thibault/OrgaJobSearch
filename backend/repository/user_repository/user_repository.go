package repository

import (
	"errors"
	"strings"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

var _ UserRepositoryInterface = &UserRepository{}

func (r *UserRepository) SaveUser(user models.User) error {
	if r.db == nil {
		return errors.New("database connection is nil")
	}

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
	email = strings.TrimSpace(email)
	result := r.db.
		Preload("Otps").
		Where("email = ?", email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) GetUserByID(ID uint) (*models.User, error) {
	if r.db == nil {
		return nil, errors.New("database connection is nil")
	}

	if ID == 0 {
		return nil, errors.New("ID cannot be zero")
	}

	var user models.User
	result := r.db.Where("id = ?", ID).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) ValidateEmail(email string) error {
	if r.db == nil {
		return errors.New("database connection is nil")
	}
	if email == "" {
		return errors.New("email cannot be empty")
	}
	return r.db.Model(&models.User{}).
		Where("email = ?", email).
		Updates(map[string]interface{}{
			"email_is_valide": true,
			"user_status":     "registered",
		}).Error

}

func (r *UserRepository) GetUserByUUID(uuid string) (*models.User, error) {
	if r.db == nil {
		return nil, errors.New("database connection is nil")
	}

	if uuid == "" {
		return nil, errors.New("uudi cannot be empty")
	}
	var user models.User
	result := r.db.
		Preload("Otps").
		Where("user_uuid = ?", uuid).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(existingUserID uint, updatedUserData models.UserProfileUpdate) error {
	if r.db == nil {
		return errors.New("database connection is nil")
	}

	result := r.db.Model(&models.User{}).Where("id = ?", existingUserID).Updates(models.User{
		LastName:  updatedUserData.UserLastName,
		FirstName: updatedUserData.UserFirstName,
	})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return result.Error
	}
	return nil
}
