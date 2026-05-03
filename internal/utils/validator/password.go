package validator

import (
	"errors"
	"unicode"
)

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool

	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasNumber = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return errors.New("password must contain at least 1 uppercase letter")
	}
	if !hasLower {
		return errors.New("must contain lowercase letter")
	}
	if !hasNumber {
		return errors.New("password must contain at least 1 number")
	}
	if !hasSpecial {
		return errors.New("password must contain at least 1 special character")
	}

	return nil
}
