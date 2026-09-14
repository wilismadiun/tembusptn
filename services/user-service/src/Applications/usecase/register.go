package usecase

import (
	"errors"
	"log"

	"github.com/wilismadiun/tembusptn/services/user-service/src/Applications/generator"
	"github.com/wilismadiun/tembusptn/services/user-service/src/Applications/security"
	domains "github.com/wilismadiun/tembusptn/services/user-service/src/Domains"
	"github.com/wilismadiun/tembusptn/services/user-service/src/Domains/entities"
)

type Register struct {
	Repo         domains.UserRepository
	Generator    generator.GeneratorId
	HashPassword security.PasswordHasher
}

var ErrNotFound = errors.New("User Not Found")

func (h *Register) Execute(user entities.User) (entities.RegisteredUser, error) {
	id := h.Generator.Generator()

	user.ID = id
	err := entities.VerifyUser(user)
	if err != nil {
		return entities.RegisteredUser{}, err
	}

	err = h.Repo.FindUserByEmail(user)
	if err == nil {
		return entities.RegisteredUser{}, errors.New("Email is already in use")
	}

	if !errors.Is(err, ErrNotFound) {
		return entities.RegisteredUser{}, err
	}

	hashedPassword, err := h.HashPassword.Hash(user.Password)
	if err != nil {
		return entities.RegisteredUser{}, err
	}

	user.Password = hashedPassword

	err = h.Repo.UserRegister(&user)
	if err != nil {
		return entities.RegisteredUser{}, err
	}

	log.Println("====================================== ini adalah user")
	log.Println(user)

	return entities.RegisteredUser{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}
