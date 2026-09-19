package usersubscriptions

import (
	"errors"
	"testing"

	"subscription-service/src/Domains/entities"
	"subscription-service/src/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAddUserSubScripitions(t *testing.T) {

	t.Run("should return error when duration is less than or equal to 0", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserSubscriptionRepository(ctrl)
		mockGenerator := mocks.NewMockGeneratorId(ctrl)

		usecase := AddUserSubScripitions{
			Repo:      mockRepo,
			Generator: mockGenerator,
		}

		result, err := usecase.Execute(
			"user-123",
			"subscription-123",
			0,
		)

		assert.Error(t, err)
		assert.Equal(t, "duration must be greater than 0", err.Error())
		assert.Equal(t, entities.UserSubscriptions{}, result)
	})

	t.Run("should return error when database failed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserSubscriptionRepository(ctrl)
		mockGenerator := mocks.NewMockGeneratorId(ctrl)

		expectedErr := errors.New("database error")

		mockGenerator.
			EXPECT().
			Generator().
			Return("user-subscription-123")

		mockRepo.
			EXPECT().
			CreateUserSubscription(gomock.Any()).
			Return(expectedErr)

		usecase := AddUserSubScripitions{
			Repo:      mockRepo,
			Generator: mockGenerator,
		}

		result, err := usecase.Execute(
			"user-123",
			"subscription-123",
			1,
		)

		assert.ErrorIs(t, err, expectedErr)
		assert.Equal(t, entities.UserSubscriptions{}, result)
	})

	t.Run("should successfully create user subscription", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserSubscriptionRepository(ctrl)
		mockGenerator := mocks.NewMockGeneratorId(ctrl)

		mockGenerator.
			EXPECT().
			Generator().
			Return("user-subscription-123")

		mockRepo.
			EXPECT().
			CreateUserSubscription(gomock.Any()).
			DoAndReturn(func(result *entities.UserSubscriptions) error {

				assert.Equal(t, "user-subscription-123", result.ID)
				assert.Equal(t, "user-123", result.UserId)
				assert.Equal(t, "subscription-123", result.SubscriptionId)
				assert.Equal(t, entities.StatusActive, result.Status)

				assert.False(t, result.StartAt.IsZero())
				assert.False(t, result.EndAt.IsZero())

				expectedEndAt := result.StartAt.AddDate(0, 1, 0)
				assert.Equal(t, expectedEndAt, result.EndAt)

				return nil
			})

		usecase := AddUserSubScripitions{
			Repo:      mockRepo,
			Generator: mockGenerator,
		}

		result, err := usecase.Execute(
			"user-123",
			"subscription-123",
			1,
		)

		assert.NoError(t, err)

		assert.Equal(t, "user-subscription-123", result.ID)
		assert.Equal(t, "user-123", result.UserId)
		assert.Equal(t, "subscription-123", result.SubscriptionId)
		assert.Equal(t, entities.StatusActive, result.Status)

		assert.False(t, result.StartAt.IsZero())
		assert.False(t, result.EndAt.IsZero())

		assert.Equal(
			t,
			result.StartAt.AddDate(0, 1, 0),
			result.EndAt,
		)
	})
}
