package repository

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"subscription-service/commons/database"
	"subscription-service/src/Applications/usecase"
	"subscription-service/src/Domains/entities"
)

var repo *SubscriptionRepository

func TestMain(m *testing.M) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Println("Peringatan: Gagal memuat .env dari root")
	}

	database.ConnectDatabase()

	repo = &SubscriptionRepository{
		DB: database.DB,
	}

	code := m.Run()

	os.Exit(code)
}

func TestFindSubscriptionByName(t *testing.T) {

	t.Run("should return subscription when found", func(t *testing.T) {

		dummySubscription := entities.Subscriptions{
			ID:    "subscription-found-123",
			Name:  "Gold Test",
			Price: 50000,
		}

		err := database.DB.Create(&dummySubscription).Error
		assert.NoError(t, err)

		err = repo.FindSubscriptionByName(dummySubscription.Name)

		assert.NoError(t, err)

		// Bersihkan data setelah test
		database.DB.
			Where("name = ?", dummySubscription.Name).
			Delete(&entities.Subscriptions{})
	})

	t.Run("should return ErrNotFound when subscription not found", func(t *testing.T) {

		err := repo.FindSubscriptionByName(
			"subscription-yang-pasti-tidak-ada",
		)

		assert.Error(t, err)
		assert.True(t, errors.Is(err, usecase.ErrNotFound))
	})
}

func TestCreateSubscriptions(t *testing.T) {

	subscription := entities.Subscriptions{
		ID:    "subscription-create-123",
		Name:  "Diamond Test",
		Price: 150000,
	}

	err := repo.CreateSubscriptions(&subscription)

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
