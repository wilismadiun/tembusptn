package usecase

import (
	"github.com/wilismadiun/tembusptn/services/user-service/src/Applications/security"
	domains "github.com/wilismadiun/tembusptn/services/user-service/src/Domains"
	"github.com/wilismadiun/tembusptn/services/user-service/src/Domains/entities"
)

type Login struct {
	UserRepo domains.UserRepository
	Token    security.AuthToken
	Hasher   security.PasswordHasher
}

func (h *Login) Execute(login entities.Login) (string, error) {
	err := entities.VerifyLogin(login)
	if err != nil {
		return "", err
	}

	existingUser, err := h.UserRepo.FindUserByEmail(login.Email)
	if err != nil {
		return "", err
	}

	err = h.Hasher.CompareHashPassword(login.Password, existingUser.Password)
	if err != nil {
		return "", err
	}

	accessToken, err := h.Token.GenerateToken(existingUser.ID)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
