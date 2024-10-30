package config

import (
	"log"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	utils "github.com/R-Thibault/OrgaJobSearch/backend/utils/hash_util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func LoadData(db *gorm.DB) error {
	roles := []models.Role{
		{RoleName: "CareerSupportManager", Description: "WIP"},
		{RoleName: "CareerCoach", Description: "WIP"},
		{RoleName: "JobSeeker", Description: "WIP"},
	}
	// Role Seeds
	for _, role := range roles {
		var existingRole models.Role

		if err := db.Where("role_name = ?", role.RoleName).First(&existingRole).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if createErr := db.Create(&role).Error; createErr != nil {
					log.Fatalf("Cannot seed role table: %v\n", err)
					return createErr
				}
				log.Printf("Role %s seeded successfully.", role.RoleName)
			} else {
				log.Fatalf("Error checking if role exists: %v\n", err)
				return err
			}
		} else {
			log.Printf("Role %s already exists, skipping seeding.", role.RoleName)
		}
	}

	//Create instance of HashingService
	hashingService := utils.NewHashingService()

	// Hash the password for the supAdmin user
	hashedPassword, err := hashingService.HashPassword("supAdmin1!")
	if err != nil {
		log.Fatalf("Error hashing password: %v\n", err)
		return err
	}

	// retreived CareerSupportManager role
	var careerSupportManagerRole models.Role
	if err := db.Where("role_name = ?", "CareerSupportManager").First(&careerSupportManagerRole).Error; err != nil {
		log.Fatalf("Error retrieving CareerSupportManager role: %v\n", err)
		return err
	}

	// retreived CareerCoach role
	var careerCoachRole models.Role
	if err := db.Where("role_name = ?", "CareerCoach").First(&careerCoachRole).Error; err != nil {
		log.Fatalf("Error retrieving CareerCoach role: %v\n", err)
		return err
	}

	supAdmin := models.User{
		FirstName:      "Maxi",
		LastName:       "SupAdmin",
		Email:          "supadmin@admin.com",
		UserUUID:       uuid.New().String(),
		HashedPassword: hashedPassword,
		EmailIsValide:  true,
		UserStatus:     "registred",
		Roles:          []models.Role{careerSupportManagerRole, careerCoachRole},
	}
	var existingSupAdmin models.User
	// Create a Admin user
	if userErr := db.Where("Last_name = ?", supAdmin.LastName).First(&existingSupAdmin).Error; userErr != nil {
		if userErr == gorm.ErrRecordNotFound {
			if createUserErr := db.Create(&supAdmin).Error; createUserErr != nil {
				log.Fatalf("Cannot seed user table: %v\n", createUserErr)
				return createUserErr
			}
			log.Printf("User %s seeded successfully.", supAdmin.LastName)

			if roleAssignErr := db.Model(&supAdmin).Association("Roles").Append(&careerSupportManagerRole); roleAssignErr != nil {
				log.Fatalf("Error assigning role to SupAdmin: %v\n", roleAssignErr)
				return roleAssignErr
			}
		} else {
			log.Fatalf("Error checking if admin user exists: %v\n", userErr)
		}
	} else {
		log.Printf("User %s already exists, skipping seeding.", supAdmin.LastName)
	}
	return nil
}
