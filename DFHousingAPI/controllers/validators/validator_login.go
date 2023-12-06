package validators

// validator for the login controller
// checks rules for email and password

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
	"fmt"
)

func CustomEmailValidation(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	fmt.Printf("checking email: %s\n", email)
	// Regular expression for validating email format
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	// Check if the email length is not more than 250 characters
	if len(email) < 6 {
		return false
	}

	if len(email) > 250 {
		return false
	}

	// Check if the email matches the regex pattern
	if !emailRegex.MatchString(email) {
		return false
	}

	return true
}

func CustomPasswordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Check if the password length is not less than 8 characters
	if len(password) < 8 {
		return false
	}

	// Check if the password length is not more than 250 characters
	if len(password) > 150 {
		return false
	}

	// Check if the password contains at least one uppercase letter
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return false
	}

	// Check if the password contains at least one lowercase letter
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return false
	}

	// Check if the password contains at least one digit
	if !strings.ContainsAny(password, "0123456789") {
		return false
	}

	// Check if the password contains at least one special character
	if !strings.ContainsAny(password, "!@#$%^&*()_+{}|:<>?") {
		return false
	}

	return true
}