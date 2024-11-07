package registrationservices

import "github.com/R-Thibault/OrgaJobSearch/backend/models"

type RegistrationServiceInterface interface {
	UserRegistration(creds models.Credentials) error
}
