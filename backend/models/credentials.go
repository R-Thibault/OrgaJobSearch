package models

// Credentials struct, for SignIn - SignUp - auth_middleware
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
