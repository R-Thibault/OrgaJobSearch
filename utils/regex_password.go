package utils

import (
	"regexp"
)

func RegexPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	// Regex to check at least one uppercase letter
	upperRegex := regexp.MustCompile(`[A-Z]`)
	// Regex to check at least one lowercase letter
	lowerRegex := regexp.MustCompile(`[a-z]`)
	// Regex to check at least one digit
	digitRegex := regexp.MustCompile(`[0-9]`)
	// Regex to check at least one special character
	specialCharRegex := regexp.MustCompile(`[@$!%*?&]`)

	// Return true only if all regexes match
	return upperRegex.MatchString(password) &&
		lowerRegex.MatchString(password) &&
		digitRegex.MatchString(password) &&
		specialCharRegex.MatchString(password)
}
