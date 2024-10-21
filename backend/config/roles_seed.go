package config

import (
	"log"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	"gorm.io/gorm"
)

func LoadData(db *gorm.DB) error {
	roles := []models.Role{
		{RoleName: "CareerSupportManager", Description: "WIP"},
		{RoleName: "CareerCoach", Description: "WIP"},
		{RoleName: "JobSeeker", Description: "WIP"},
	}

	for _, role := range roles {
		var existingRole models.Role

		if err := db.Where("role_name = ?", role.RoleName).First(&existingRole).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if createErr := db.Create(&role).Error; err != nil {
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
	return nil
}
