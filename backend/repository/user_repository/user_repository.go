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
	return r.db.Model(&models.User{}).
		Where("email = ?", email).
		Updates(map[string]interface{}{
			"email_is_valide": true,
			"user_status":     "registered",
		}).Error

}

func (r *UserRepository) PreRegisterJobSeeker(user models.User) (*models.User, error) {
	if user.Email == "" {
		return nil, errors.New("email cannot be empty")
	}
	user.UserStatus = "pre-registred"
	// Save user in DB
	result := r.db.Create(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) GetUserByUUID(uuid string) (*models.User, error) {
	if uuid == "" {
		return nil, errors.New("uudi cannot be empty")
	}
	var user models.User
	result := r.db.Unscoped().Where("user_uuid = ?", uuid).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) UpdateJobSeeker(savedUser models.User) error {
	result := r.db.Save(&models.User{
		Model: gorm.Model{
			ID: savedUser.ID,
		},
		Email:          savedUser.Email,
		HashedPassword: savedUser.HashedPassword,
		UserUUID:       savedUser.UserUUID,
		UserStatus:     "registred",
		EmailIsValide:  true})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return result.Error
	}
	return nil
}
