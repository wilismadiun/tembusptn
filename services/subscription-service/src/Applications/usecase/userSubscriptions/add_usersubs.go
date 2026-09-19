package usersubscriptions

import (
	"errors"
	"time"

	"subscription-service/src/Applications/generator"
	domains "subscription-service/src/Domains"
	"subscription-service/src/Domains/entities"
)

type AddUserSubScripitions struct {
	Repo      domains.UserSubscriptionRepository
	Generator generator.GeneratorId
}

func (h *AddUserSubScripitions) Execute(userId, subscriptionId string, duration int) (entities.UserSubscriptions, error) {

	if duration <= 0 {
		return entities.UserSubscriptions{}, errors.New("duration must be greater than 0")
	}

	id := h.Generator.Generator()

	now := time.Now()

	startAt := now
	endAt := now.AddDate(0, duration, 0)
	status := entities.StatusActive

	userSub := entities.UserSubscriptions{
		ID:             id,
		UserId:         userId,
		SubscriptionId: subscriptionId,
		StartAt:        startAt,
		EndAt:          endAt,
		Status:         status,
	}

	err := h.Repo.CreateUserSubscription(&userSub)
	if err != nil {
		return entities.UserSubscriptions{}, err
	}

	return userSub, nil
}
