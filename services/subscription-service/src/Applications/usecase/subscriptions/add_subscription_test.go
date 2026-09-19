package subscriptions

import (
	"errors"
	"testing"

	"subscription-service/src/Domains/entities"
	"subscription-service/src/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAddSubscription(t *testing.T) {
	t.Run("should return error when name is empty", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockSubscriptionRepository(ctrl)
		generator := mocks.NewMockGeneratorId(ctrl)

		handler := AddSubscription{
			Repo:      repo,
			Generator: generator,
		}

		subs := entities.Subscriptions{
			Name:  "",
			Price: 50000,
		}

		id, err := handler.Execute(subs)

		assert.Empty(t, id)
		assert.Error(t, err)
		assert.Equal(t, "name is required", err.Error())
	})

	t.Run("should return error when price is empty", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockSubscriptionRepository(ctrl)
		generator := mocks.NewMockGeneratorId(ctrl)

		handler := AddSubscription{
			Repo:      repo,
			Generator: generator,
		}

		subs := entities.Subscriptions{
			Name:  "Gold",
			Price: 0,
		}

		id, err := handler.Execute(subs)

		assert.Empty(t, id)
		assert.Error(t, err)
		assert.Equal(t, "price is required", err.Error())
	})

	t.Run("should return error when subscription already exists", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockSubscriptionRepository(ctrl)
		generator := mocks.NewMockGeneratorId(ctrl)

		handler := AddSubscription{
			Repo:      repo,
			Generator: generator,
		}

		subs := entities.Subscriptions{
			Name:  "Gold",
			Price: 50000,
		}

		generator.EXPECT().
			Generator().
			Return("subscription-123")

		repo.EXPECT().
			FindSubscriptionByName("Gold").
			Return(nil)

		id, err := handler.Execute(subs)

		assert.Empty(t, id)
		assert.Error(t, err)
		assert.Equal(t, "Subscription is already in use", err.Error())
	})

	t.Run("should return error when repository failed to create subscription", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockSubscriptionRepository(ctrl)
		generator := mocks.NewMockGeneratorId(ctrl)

		handler := AddSubscription{
			Repo:      repo,
			Generator: generator,
		}

		subs := entities.Subscriptions{
			Name:  "Gold",
			Price: 50000,
		}

		generator.EXPECT().
			Generator().
			Return("subscription-123")

		repo.EXPECT().
			FindSubscriptionByName("Gold").
			Return(ErrNotFound)

		repo.EXPECT().
			CreateSubscriptions(gomock.Any()).
			Return(errors.New("failed to create subscription"))

		id, err := handler.Execute(subs)

		assert.Empty(t, id)
		assert.Error(t, err)
		assert.Equal(t, "failed to create subscription", err.Error())
	})

	t.Run("should return subscription id when success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockSubscriptionRepository(ctrl)
		generator := mocks.NewMockGeneratorId(ctrl)

		handler := AddSubscription{
			Repo:      repo,
			Generator: generator,
		}

		subs := entities.Subscriptions{
			Name:  "Gold",
			Price: 50000,
		}

		generator.EXPECT().
			Generator().
			Return("subscription-123")

		repo.EXPECT().
			FindSubscriptionByName("Gold").
			Return(ErrNotFound)

		repo.EXPECT().
			CreateSubscriptions(gomock.Any()).
			DoAndReturn(func(subs *entities.Subscriptions) error {
				assert.Equal(t, "subscription-123", subs.ID)
				assert.Equal(t, "Gold", subs.Name)
				assert.Equal(t, int64(50000), subs.Price)

				return nil
			})

		id, err := handler.Execute(subs)

		assert.NoError(t, err)
		assert.Equal(t, "subscription-123", id)
	})
}
