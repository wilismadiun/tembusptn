package infrastructures

import (
	"github.com/wilismadiun/tembusptn/services/user-service/src/Applications/usecase"
	"github.com/wilismadiun/tembusptn/services/user-service/src/Infrastructures/generator"
	"github.com/wilismadiun/tembusptn/services/user-service/src/Infrastructures/repository"
	"github.com/wilismadiun/tembusptn/services/user-service/src/Infrastructures/security"
	"github.com/wilismadiun/tembusptn/services/user-service/src/Interfaces/http"
	"gorm.io/gorm"
)

func Container(db *gorm.DB) *http.Handler {
	userRepoImplement := repository.UserRepository{DB: db}
	roleRepoImplement := repository.RoleRepository{DB: db}
	hasherImplement := security.HashPasswordBcrypt{}
	generatorImplement := generator.GeneratorUUID{}
	tokenImplement := security.AuthenticationTokenJWT{}

	registerHandler := usecase.Register{
		Repo:         &userRepoImplement,
		Generator:    &generatorImplement,
		HashPassword: &hasherImplement,
	}

	loginHandler := usecase.Login{
		UserRepo: &userRepoImplement,
		Token:    &tokenImplement,
		Hasher:   &hasherImplement,
		RoleRepo: &roleRepoImplement,
	}

	return &http.Handler{
		RegisterHandler: &registerHandler,
		Loginhandler:    &loginHandler,
	}
}
