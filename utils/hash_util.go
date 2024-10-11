package utils

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

// HashPassword generates a hashed password using the Argon2id algorithm.
func HashPassword(password string) (string, error) {
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
