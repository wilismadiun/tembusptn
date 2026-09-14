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
	repoImplement := repository.UserRepository{DB: db}
	hasherImplement := security.HashPasswordBcrypt{}
	generatorImplement := generator.GeneratorUUID{}

	registerHandler := usecase.Register{
		Repo:         &repoImplement,
		Generator:    &generatorImplement,
		HashPassword: &hasherImplement,
	}

	return &http.Handler{
		RegisterHandler: &registerHandler,
	}
}
