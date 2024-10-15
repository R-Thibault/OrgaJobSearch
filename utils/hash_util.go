package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type HashingService struct{}

func NewHashingService() *HashingService {

	return &HashingService{}

}

var _ HashingServiceInterface = &HashingService{}

// HashPassword generates a hashed password using the Argon2id algorithm.
func (h *HashingService) HashPassword(password string) (string, error) {
	// Generate a random salt of 16 bytes
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// Generate the hash using Argon2id with the given parameters
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	// Encode the hash and salt to base64
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)

	// Return the encoded salt and hash, separated by a colon
	return encodedSalt + ":" + encodedHash, nil
}

func (h *HashingService) CompareHashPassword(password, storedHash string) (bool, error) {
	// Split the stored hash into salt and hash components
	saltAndHash := strings.Split(storedHash, ":")
	if len(saltAndHash) != 2 {
		return false, errors.New("invalid stored hash format")
	}

	salt, err := base64.RawStdEncoding.DecodeString(saltAndHash[0])
	if err != nil {
		return false, fmt.Errorf("failed to decode salt: %w", err)
	}

	hash, err := base64.RawStdEncoding.DecodeString(saltAndHash[1])
	if err != nil {
		return false, fmt.Errorf("failed to decode hash: %w", err)
	}

	// Generate a new hash using the provided password and the extracted salt
	newHash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	// Compare the newly generated hash with the stored hash
	if subtle.ConstantTimeCompare(newHash, hash) == 1 {
		return true, nil
	}

	return false, nil
}
