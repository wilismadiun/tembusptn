package subscriptions

import (
	domains "subscription-service/src/Domains"
	"subscription-service/src/Domains/entities"
)

type GetAllSubscriptions struct {
	Repo domains.SubscriptionRepository
}

func (h *GetAllSubscriptions) Execute() ([]entities.Subscriptions, error) {
	result, err := h.Repo.GetAllSubscriptions()
	if err != nil {
		return []entities.Subscriptions{}, err
	}

	return result, nil
}
