package infrastructures

import (
	"subscription-service/src/Applications/usecase/subscriptions"
	usersubscriptions "subscription-service/src/Applications/usecase/userSubscriptions"
	"subscription-service/src/Infrastructures/generator"
	"subscription-service/src/Infrastructures/repository"
	"subscription-service/src/Interfaces/http"

	"gorm.io/gorm"
)

func SubscriptionsContainer(db *gorm.DB) *http.SubscriptionsHandler {
	repo := repository.SubscriptionRepository{DB: db}
	generatorId := generator.GeneratorUUID{}

	addSubscriptionHandler := subscriptions.AddSubscription{
		Repo:      &repo,
		Generator: &generatorId,
	}

	getAllSubscriptionHandler := subscriptions.GetAllSubscriptions{
		Repo: &repo,
	}

	return &http.SubscriptionsHandler{
		AddSubscriptionHandler:    &addSubscriptionHandler,
		GetAllSubscriptionHandler: &getAllSubscriptionHandler,
	}
}

func UserSubsContsiner(db *gorm.DB) *http.UserSubsHandler {
	repo := repository.UserSubscriptionRepository{DB: db}
	generatorId := generator.GeneratorUUID{}

	addUserSubscriptionHandler := usersubscriptions.AddUserSubScripitions{
		Repo:      &repo,
		Generator: &generatorId,
	}

	return &http.UserSubsHandler{
		AddUserSubscriptionHandler: &addUserSubscriptionHandler,
	}
}
