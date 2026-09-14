package domains

import "github.com/wilismadiun/tembusptn/services/user-service/src/Domains/entities"

type UserRepository interface {
	UserRegister(user *entities.User) error
	FindUserByEmail(email string) (entities.User, error)
}
