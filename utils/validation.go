package utils

import (
	"errors"
	"regexp"
)

func ValidateUsername(username string) error {
	matched, err := regexp.MatchString("^[a-zA-Z0-9]+$", username)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("username must be alphanumeric")
	}
	return nil
}

func ValidateEmail(email string) error {
	matched, err := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("invalid email format")
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	matched, err := regexp.MatchString(`[0-9]`, password)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("password must contain at least one number")
	}
	matched, err = regexp.MatchString(`[a-zA-Z]`, password)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("password must contain at least one letter")
	}
	return nil
}
