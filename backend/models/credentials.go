package models

// Credentials struct, for SignIn - SignUp - auth_middleware
type Credentials struct {
	Email       string `json:"email"`
	LastName    string `json:"lastName"`
	FirstName   string `json:"firstName"`
	Password    string `json:"password"`
	TokenString string `json:"tokenString"`
}
