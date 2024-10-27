package invitationservices

import "github.com/R-Thibault/OrgaJobSearch/backend/models"

type InvitationServiceInterface interface {
	VerifyPersonnalInvitationTokenData(token models.JWTToken) (email string, err error)
	VerifyGlobalInvitationTokenData(token models.JWTToken) error
}
