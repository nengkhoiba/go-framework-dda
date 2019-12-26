package utils

import (
	"errors"
	"regexp"
)

var (
	errPasswordEmpty    = errors.New("Password field is empty")
	errUsernameEmpty    = errors.New("Username field is empty")
	errEmailEmpty       = errors.New("Email field is empty")
	errUsernameInvalid  = errors.New("Invalid username, allowed characters: A-Z, a-z, 0-9, . and _")
	errInvalidEmail     = errors.New("Invalid email")
	errPasswordTooShort = errors.New("Password too short, should be more than 5 characters")
	errPasswordTooLong  = errors.New("Password too long, disallowed to avoid DOS")
	errUsernameTooShort = errors.New("Username is too short, should be more than 3 characters")
	errUsernameTooLong  = errors.New("Username is too long, should be less than 30 characters")
	errEmailTooLong     = errors.New("Email too long, should be less than 1024 characters")
	errNameTooLong      = errors.New("Name too long, should be less than 1024 characters")
	errNameEmpty        = errors.New("Name field is empty")
)

// UsernameRegex function to validate username
var UsernameRegex = regexp.MustCompile("^[A-Za-z0-9_.]+$")
var UsernameInverseRegex = regexp.MustCompile("[^A-Za-z0-9_.]+")

// EmailRegex function to validate email
var EmailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// ValidateUsername checks if given username is valid
// Requirements 3 to 30 chars
// Latin alphabets , 0-9 and underscores
// Case insensitive
func ValidateUsername(username string) error {
	if len(username) == 0 {
		return errUsernameEmpty
	}
	if len(username) < 3 {
		return errUsernameTooShort
	}

	if len(username) > 30 {
		return errUsernameTooLong
	}
	if !UsernameRegex.MatchString(username) {
		return errUsernameInvalid
	}
	return nil
}

// ValidateFullName for full name validation
func ValidateFullName(fullname string) error {
	if len(fullname) == 0 {
		return errNameEmpty
	}
	if len(fullname) > 1024 {
		return errNameTooLong
	}
	return nil
}

// ValidateEmail for email validation
func ValidateEmail(email string) error {
	if !EmailRegex.MatchString(email) {
		return errInvalidEmail
	}
	if len(email) > 1024 {
		return errEmailTooLong
	}
	return nil
}

// ValidatePassword for password input validation
func ValidatePassword(password string) error {
	if len(password) < 5 {
		return errPasswordTooShort
	}
	if len(password) > 1024 {
		return errPasswordTooLong
	}
	return nil
}

func IsValidEmail(email string) bool {
	if !EmailRegex.MatchString(email) {
		return false
	}
	if len(email) > 1024 {
		return false
	}
	return true
}

func IsValidUsername(username string) bool {
	if len(username) == 0 {
		return false
	}
	if len(username) < 3 {
		return false
	}

	if len(username) > 30 {
		return false
	}
	if !UsernameRegex.MatchString(username) {
		return false
	}
	return true
}

// IsValidTopicDetails

// IsValidCategoryName

// IsValidPhone
