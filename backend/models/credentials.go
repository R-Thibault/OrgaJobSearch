package models

// Credentials struct, for SignIn - SignUp - auth_middleware
type Credentials struct {
	Email           string `json:"email"`
	LastName        string `json:"lastName"`
	FirstName       string `json:"firstName"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type ResetPasswordCredentials struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	TokenString     string `json:"tokenString"`
}
