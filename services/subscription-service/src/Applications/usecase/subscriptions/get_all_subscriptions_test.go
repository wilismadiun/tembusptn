package subscriptions

import (
	"errors"
	"testing"

	"subscription-service/src/Domains/entities"
	"subscription-service/src/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAllSubscriptions(t *testing.T) {
	t.Run("should return error when repository failed", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockRepo := mocks.NewMockSubscriptionRepository(ctrl)

		expectedErr := errors.New("failed to get subscriptions")

		mockRepo.
			EXPECT().
			GetAllSubscriptions().
			Return([]entities.Subscriptions{}, expectedErr)

		getAllSubscriptions := GetAllSubscriptions{
			Repo: mockRepo,
		}

		result, err := getAllSubscriptions.Execute()

		assert.ErrorIs(t, err, expectedErr)
		assert.Equal(t, []entities.Subscriptions{}, result)
	})

	t.Run("should return all subscriptions successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockRepo := mocks.NewMockSubscriptionRepository(ctrl)

		expected := []entities.Subscriptions{
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

		mockRepo.
			EXPECT().
			GetAllSubscriptions().
			Return(expected, nil)

		getAllSubscriptions := GetAllSubscriptions{
			Repo: mockRepo,
		}

		result, err := getAllSubscriptions.Execute()

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}
