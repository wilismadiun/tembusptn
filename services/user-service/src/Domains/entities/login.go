package entities

import (
	"errors"
	"net/mail"
)

func VerifyLogin(login Login) error {
	if login.Email == "" {
		return errors.New("email is required")
	}

	_, err := mail.ParseAddress(login.Email)
	if err != nil {
		return errors.New("invalid email format")
	}

	if login.Password == "" {
		return errors.New("password is required")
	}

	if len(login.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}
