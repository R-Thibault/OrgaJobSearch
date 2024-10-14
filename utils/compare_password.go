package utils

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

func CompareHashPassword(password, storedHash string) (bool, error) {
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
