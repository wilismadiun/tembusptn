package repository

import (
	"errors"

	"subscription-service/src/Applications/usecase"
	"subscription-service/src/Domains/entities"

	"gorm.io/gorm"
)

type SubscriptionRepository struct {
	DB *gorm.DB
}

func (r *SubscriptionRepository) FindSubscriptionByName(name string) error {
	var subs entities.Subscriptions

	err := r.DB.Where("name = ?", name).First(&subs).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return usecase.ErrNotFound
		} else {
			return err
		}
	}

	return nil
}

func (r *SubscriptionRepository) CreateSubscriptions(subs *entities.Subscriptions) error {
	return r.DB.Create(subs).Error
}

func (r *SubscriptionRepository) GetAllSubscriptions() ([]entities.Subscriptions, error) {
	var subscriptions []entities.Subscriptions

	err := r.DB.Find(&subscriptions).Error
	if err != nil {
		return []entities.Subscriptions{}, err
	}

	return subscriptions, nil
}
