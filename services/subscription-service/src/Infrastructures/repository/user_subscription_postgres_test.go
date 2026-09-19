package repository

import (
	"testing"
	"time"

	"subscription-service/commons/database"
	"subscription-service/src/Domains/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func helperDatabase() {
	database.DB.Exec("DELETE FROM user_subscriptions")
	database.DB.Exec("DELETE FROM users")
	database.DB.Exec("DELETE FROM subscriptions")

	database.DB.Exec(`
		INSERT INTO roles (id, name)
		VALUES
			(1, 'admin'),
			(2, 'student'),
			(3, 'teacher')
		ON CONFLICT (id) DO NOTHING
	`)

	database.DB.Exec(`
		INSERT INTO users (
			id,
			role_id,
			email,
			name,
			phone,
			password
		)
		VALUES (
			'user-123',
			2,
			'user@gmail.com',
			'Jaya',
			'081234567890',
			'password123'
		)
	`)

	database.DB.Exec(`
		INSERT INTO subscriptions (
			id,
			name,
			price
		)
		VALUES (
			'subscription-123',
			'Gold',
			50000
		)
	`)
}

func TestCreateUserSubscription(t *testing.T) {
	helperDatabase()

	userSub := entities.UserSubscriptions{
		ID:             "user-subscription-123",
		UserId:         "user-123",
		SubscriptionId: "subscription-123",
		StartAt:        time.Now(),
		EndAt:          time.Now().AddDate(0, 1, 0),
		Status:         entities.StatusActive,
	}

	err := userSubsRepo.CreateUserSubscription(&userSub)

	require.NoError(t, err)

	var result entities.UserSubscriptions

	err = database.DB.
		Where("id = ?", userSub.ID).
		First(&result).Error

	require.NoError(t, err)

	assert.Equal(t, userSub.ID, result.ID)
	assert.Equal(t, userSub.UserId, result.UserId)
	assert.Equal(t, userSub.SubscriptionId, result.SubscriptionId)
	assert.Equal(t, userSub.Status, result.Status)

	assert.WithinDuration(
		t,
		userSub.StartAt.UTC(),
		result.StartAt.UTC(),
		time.Second,
	)

	assert.WithinDuration(
		t,
		userSub.EndAt.UTC(),
		result.EndAt.UTC(),
		time.Second,
	)

	assert.Equal(t, userSub.Status, result.Status)
}
