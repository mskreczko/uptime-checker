package pkg

import (
	"regexp"
	"strings"
)

func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    return emailRegex.MatchString(email)
}

func ValidatePhoneNumber(phoneNumber string) bool {
    re := regexp.MustCompile(`^\+[1-9]\d{1,14}$`) //E164 standard
    phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")
	return re.MatchString(phoneNumber)
}