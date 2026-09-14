package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilismadiun/tembusptn/services/user-service/src/Domains/entities"
	"github.com/wilismadiun/tembusptn/services/user-service/src/mocks"
	"go.uber.org/mock/gomock"
)

func Test_Login(t *testing.T) {
	t.Run("should return error when email is empty", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockRepo := mocks.NewMockUserRepository(ctrl)
		mockToken := mocks.NewMockAuthToken(ctrl)
		mockHasher := mocks.NewMockPasswordHasher(ctrl)

		login := Login{
			UserRepo: mockRepo,
			Token:    mockToken,
			Hasher:   mockHasher,
		}

		input := entities.Login{
			Email:    "",
			Password: "password123",
		}

		result, err := login.Execute(input)

		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("should return error when email is not registered", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockRepo := mocks.NewMockUserRepository(ctrl)
		mockToken := mocks.NewMockAuthToken(ctrl)
		mockHasher := mocks.NewMockPasswordHasher(ctrl)

		mockRepo.
			EXPECT().
			FindUserByEmail("user@example.com").
			Return(entities.User{}, assert.AnError)

		login := Login{
			UserRepo: mockRepo,
			Token:    mockToken,
			Hasher:   mockHasher,
		}

		input := entities.Login{
			Email:    "user@example.com",
			Password: "password123",
		}

		result, err := login.Execute(input)

		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("should return error when password does not match", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockRepo := mocks.NewMockUserRepository(ctrl)
		mockToken := mocks.NewMockAuthToken(ctrl)
		mockHasher := mocks.NewMockPasswordHasher(ctrl)

		existingUser := entities.User{
			ID:       "user-123",
			Email:    "user@example.com",
			Password: "hashed-password",
		}

		mockRepo.
			EXPECT().
			FindUserByEmail("user@example.com").
			Return(existingUser, nil)

		mockHasher.
			EXPECT().
			CompareHashPassword("wrongpassword", "hashed-password").
			Return(assert.AnError)

		login := Login{
			UserRepo: mockRepo,
			Token:    mockToken,
			Hasher:   mockHasher,
		}

		input := entities.Login{
			Email:    "user@example.com",
			Password: "wrongpassword",
		}

		result, err := login.Execute(input)

		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("should return access token when login is successful", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockRepo := mocks.NewMockUserRepository(ctrl)
		mockToken := mocks.NewMockAuthToken(ctrl)
		mockHasher := mocks.NewMockPasswordHasher(ctrl)

		existingUser := entities.User{
			ID:       "user-123",
			Email:    "user@example.com",
			Password: "hashed-password",
		}

		mockRepo.
			EXPECT().
			FindUserByEmail("user@example.com").
			Return(existingUser, nil)

		mockHasher.
			EXPECT().
			CompareHashPassword("password123", "hashed-password").
			Return(nil)

		mockToken.
			EXPECT().
			GenerateToken("user-123").
			Return("access-token-123", nil)

		login := Login{
			UserRepo: mockRepo,
			Token:    mockToken,
			Hasher:   mockHasher,
		}

		input := entities.Login{
			Email:    "user@example.com",
			Password: "password123",
		}

		result, err := login.Execute(input)

		assert.NoError(t, err)
		assert.Equal(t, "access-token-123", result)
	})
}
