package repository

import (
	"subscription-service/src/Domains/entities"

	"gorm.io/gorm"
)

type UserSubscriptionRepository struct {
	DB *gorm.DB
}

func (r *UserSubscriptionRepository) CreateUserSubscription(userSub *entities.UserSubscriptions) error {
	err := r.DB.Create(userSub).Error
	if err != nil {
		return err
	}

	return nil
}
