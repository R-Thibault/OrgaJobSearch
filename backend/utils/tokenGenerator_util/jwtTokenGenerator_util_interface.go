package utils

import "time"

type JWTTokenGeneratorUtilInterface interface {
	// GenerateJWTToken generates a JWT token with the specified type, body, and expiration time.
	// It uses the JWT key from the configuration to sign the token.
	//
	// Parameters:
	//   - tokenType: A string representing the type of the token.
	//   - body: A string representing the body of the token.
	//   - expirationTime: A time.Time value representing the expiration time of the token.
	//
	// Returns:
	//   - JWTToken: A string representing the generated JWT token.
	//   - err: An error if the token generation or signing fails.
	GenerateJWTToken(tokenType string, body string, expirationTime time.Time) (JWTToken string, err error)
}
