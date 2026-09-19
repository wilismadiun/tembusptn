package repository

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"subscription-service/commons/database"
	"subscription-service/src/Applications/usecase/subscriptions"
	"subscription-service/src/Domains/entities"
)

func TestFindSubscriptionByName(t *testing.T) {

	t.Run("should return subscription when found", func(t *testing.T) {

		dummySubscription := entities.Subscriptions{
			ID:    "subscription-found-123",
			Name:  "Gold Test",
			Price: 50000,
		}

		err := database.DB.Create(&dummySubscription).Error
		assert.NoError(t, err)

		err = subscriptionRepo.FindSubscriptionByName(dummySubscription.Name)

		assert.NoError(t, err)

		// Bersihkan data setelah test
		database.DB.
			Where("name = ?", dummySubscription.Name).
			Delete(&entities.Subscriptions{})
	})

	t.Run("should return ErrNotFound when subscription not found", func(t *testing.T) {

		err := subscriptionRepo.FindSubscriptionByName(
			"subscription-yang-pasti-tidak-ada",
		)

		assert.Error(t, err)
		assert.True(t, errors.Is(err, subscriptions.ErrNotFound))
	})
}

func TestCreateSubscriptions(t *testing.T) {

	subscription := entities.Subscriptions{
		ID:    "subscription-create-123",
		Name:  "Diamond Test",
		Price: 150000,
	}

	err := subscriptionRepo.CreateSubscriptions(&subscription)

	assert.NoError(t, err)

	var result entities.Subscriptions

	err = database.DB.
		Where("id = ?", subscription.ID).
		First(&result).
		Error

	assert.NoError(t, err)

	assert.Equal(t, subscription.ID, result.ID)
	assert.Equal(t, subscription.Name, result.Name)
	assert.Equal(t, subscription.Price, result.Price)

	t.Logf("Subscription berhasil disimpan: %+v", result)

	database.DB.
		Where("id = ?", subscription.ID).
		Delete(&entities.Subscriptions{})
}

func TestGetAllSubscriptions(t *testing.T) {
	database.DB.Exec("DELETE FROM subscriptions")

	subscriptions := []entities.Subscriptions{
		{
			ID:    "subscription-1",
			Name:  "Gold",
			Price: 50000,
		},
		{
			ID:    "subscription-2",
			Name:  "Diamond",
			Price: 150000,
		},
	}

	err := database.DB.Create(&subscriptions).Error
	assert.NoError(t, err)

	// Execute
	result, err := subscriptionRepo.GetAllSubscriptions()

	// Assert
	assert.NoError(t, err)
	assert.Len(t, result, 2)

	assert.Equal(t, "subscription-1", result[0].ID)
	assert.Equal(t, "Gold", result[0].Name)
	assert.Equal(t, int64(50000), result[0].Price)

	assert.Equal(t, "subscription-2", result[1].ID)
	assert.Equal(t, "Diamond", result[1].Name)
	assert.Equal(t, int64(150000), result[1].Price)

	// Cleanup
	database.DB.Exec("DELETE FROM subscriptions")
}
