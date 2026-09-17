package entities

import (
	"errors"
	"regexp"
	"strings"
)

func VerifyUser(user User) error {
	if strings.TrimSpace(user.ID) == "" {
		return errors.New("id is required")
	}

	if user.RoleId == 0 {
		return errors.New("role id is required")
	}

	if strings.TrimSpace(user.Email) == "" {
		return errors.New("email is required")
	}

	if strings.TrimSpace(user.Name) == "" {
		return errors.New("name is required")
	}

	if strings.TrimSpace(user.Phone) == "" {
		return errors.New("telephone is required")
	}

	if strings.TrimSpace(user.Password) == "" {
		return errors.New("password is required")
	}

	emailRegex := regexp.MustCompile(
		`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
	)

	if !emailRegex.MatchString(user.Email) {
		return errors.New("invalid email format")
	}

	telephoneRegex := regexp.MustCompile(`^\+?[0-9]{10,15}$`)

	if !telephoneRegex.MatchString(user.Phone) {
		return errors.New("telephone must contain 10 to 15 digits")
	}

	if len(user.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}
