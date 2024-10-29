package registrationservices

import "github.com/R-Thibault/OrgaJobSearch/backend/models"

type RegistrationServiceInterface interface {
	RegisterCareerCoach(creds models.Credentials) error
	PreRegisterJobSeeker(email string, careerSuportID *uint) (*models.User, error)
	JobSeekerRegistration(tokenBody string, creds models.Credentials) error
}
