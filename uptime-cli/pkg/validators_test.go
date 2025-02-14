package pkg

import (
	"testing"
)



func TestEmailValidation(t *testing.T) {
	emailTestCases := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},
		{"user.name+tag@example.co.uk", true},
		{"user@subdomain.example.com", true},
		{"", false},
		{"invalid.email", false},
		{"@missingusername.com", false},
		{"user@.com", false},
		{"user@domain", false},
		{"user@domain.", false},
	}

	for _, tc := range emailTestCases {
		t.Run(tc.email, func(t *testing.T) {
			result := ValidateEmail(tc.email)
			if result != tc.expected {
				t.Errorf("ValidateEmail(%q) = %v; want %v", tc.email, result, tc.expected)
			}
		})
	}
}

func TestPhoneNumberValidation(t *testing.T) {
	phoneTestCases := []struct {
		phoneNumber string
		expected bool
	}{
		{"+14155552671", true},
		{"+12345", true},
		{"+14151234567", true},
		{"14155552671", false},
		{"", false},
		{"+", false},
	}

	for _, tc := range phoneTestCases {
		t.Run(tc.phoneNumber, func(t *testing.T) {
			result := ValidatePhoneNumber(tc.phoneNumber)
			if result != tc.expected {
				t.Errorf("ValidatePhoneNumber(%q) = %v; want %v", tc.phoneNumber, result, tc.expected)
			}
		})
	}
}