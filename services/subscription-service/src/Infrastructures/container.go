package infrastructures

import (
	"subscription-service/src/Applications/usecase"
	"subscription-service/src/Infrastructures/generator"
	"subscription-service/src/Infrastructures/repository"
	"subscription-service/src/Interfaces/http"

	"gorm.io/gorm"
)

func Container(db *gorm.DB) *http.Handler {
	repo := repository.SubscriptionRepository{DB: db}
	generatorId := generator.GeneratorUUID{}

	addSubscriptionHandler := usecase.AddSubscription{
		Repo:      &repo,
		Generator: &generatorId,
	}

	getAllSubscriptionHandler := usecase.GetAllSubscriptions{
		Repo: &repo,
	}

	return &http.Handler{
		AddSubscriptionHandler:    &addSubscriptionHandler,
		GetAllSubscriptionHandler: &getAllSubscriptionHandler,
	}
}
