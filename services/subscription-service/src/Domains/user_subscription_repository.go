package domains

import "subscription-service/src/Domains/entities"

type UserSubscriptionRepository interface {
	CreateUserSubscription(userSub *entities.UserSubscriptions) error
}
