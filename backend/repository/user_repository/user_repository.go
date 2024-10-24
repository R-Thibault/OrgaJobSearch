package repository

import (
	"errors"
	"log"
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
	result := r.db.Unscoped().Where("email = ?", email).First(&user)
	if result.Error == nil {
		log.Printf("User found: %+v\n", user)
	} else {
		log.Printf("User not found or other error: %v\n", result.Error)
	}

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
	result := r.db.Unscoped().Where("id = ?", ID).First(&user)
	if result.Error == nil {
		log.Printf("User found: %+v\n", user)
	} else {
		log.Printf("User not found or other error: %v\n", result.Error)
	}

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}
	return r.db.Model(&models.User{}).Where("email = ?", email).Update("email_is_valide", true).Error

}

func (r *UserRepository) PreRegisterUser(user models.User) error {
	if user.Email == "" {
		return errors.New("email cannot be empty")
	}
	// Save user in DB
	return r.db.Create(&user).Error
}
