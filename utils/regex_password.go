package utils

import (
	"log"
	"regexp"
)

func RegexPassword(password string) bool {
	isMatch, err := regexp.Match("^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$", []byte(password))
	if err != nil {
		log.Printf("NOK : %v\n", err)
		return false
	}
	return true
}
