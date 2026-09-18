package domains

import "subscription-service/src/Domains/entities"

type SubscriptionRepository interface {
	FindSubscriptionByName(name string) error
	CreateSubscriptions(subs *entities.Subscriptions) error
	GetAllSubscriptions() ([]entities.Subscriptions, error)
}
