package utils

type HashingServiceInterface interface {
	HashPassword(password string) (string, error)
	CompareHashPassword(password, hashedPassword string) (bool, error)
}
