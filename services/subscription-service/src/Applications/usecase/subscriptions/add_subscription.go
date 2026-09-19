package subscriptions

import (
	"errors"
	"subscription-service/src/Applications/generator"
	domains "subscription-service/src/Domains"
	"subscription-service/src/Domains/entities"
)

type AddSubscription struct {
	Repo      domains.SubscriptionRepository
	Generator generator.GeneratorId
}

var ErrNotFound = errors.New("Subscription Not Found")

func (h *AddSubscription) Execute(subs entities.Subscriptions) (string, error) {
	err := entities.VerifySubscription(subs)
	if err != nil {
		return "", err
	}

	subs.ID = h.Generator.Generator()

	err = h.Repo.FindSubscriptionByName(subs.Name)
	if err == nil {
		return "", errors.New("Subscription is already in use")
	}

	if !errors.Is(err, ErrNotFound) {
		return "", err
	}

	err = h.Repo.CreateSubscriptions(&subs)
	if err != nil {
		return "", err
	}

	return subs.ID, nil
}
