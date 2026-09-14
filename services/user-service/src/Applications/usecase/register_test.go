package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilismadiun/tembusptn/services/user-service/src/Domains/entities"
	"github.com/wilismadiun/tembusptn/services/user-service/src/mocks"
	"go.uber.org/mock/gomock"
)

func Test_register(t *testing.T) {
	t.Run("should return error when verify user fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockRepo := mocks.NewMockUserRepository(ctrl)
		mockGenerator := mocks.NewMockGeneratorId(ctrl)
		mockHasher := mocks.NewMockPasswordHasher(ctrl)

		mockGenerator.
			EXPECT().
			Generator().
			Return("user-123")

		register := Register{
			Repo:         mockRepo,
			Generator:    mockGenerator,
			HashPassword: mockHasher,
		}

		user := entities.User{}

		result, err := register.Execute(user)

		assert.Error(t, err)
		assert.Equal(t, entities.RegisteredUser{}, result)
	})

	t.Run("should return error when email already exists", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockRepo := mocks.NewMockUserRepository(ctrl)
		mockGenerator := mocks.NewMockGeneratorId(ctrl)
		mockHasher := mocks.NewMockPasswordHasher(ctrl)

		expectedErr := errors.New("email already exists")

		mockGenerator.
			EXPECT().
			Generator().
			Return("user-123")

		mockRepo.
			EXPECT().
			FindUserByEmail(gomock.Any()).
			Return(expectedErr)

		register := Register{
			Repo:         mockRepo,
			Generator:    mockGenerator,
			HashPassword: mockHasher,
		}

		user := entities.User{
			Email:    "test@example.com",
			Name:     "John",
			Phone:    "08123456789",
			Password: "password123",
		}

		result, err := register.Execute(user)

		assert.ErrorIs(t, err, expectedErr)
		assert.Equal(t, entities.RegisteredUser{}, result)
	})

	t.Run("should return error when hash password fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockRepo := mocks.NewMockUserRepository(ctrl)
		mockGenerator := mocks.NewMockGeneratorId(ctrl)
		mockHasher := mocks.NewMockPasswordHasher(ctrl)

		expectedErr := errors.New("failed to hash password")

		mockGenerator.
			EXPECT().
			Generator().
			Return("user-123")

		mockRepo.
			EXPECT().
			FindUserByEmail(gomock.Any()).
			Return(ErrNotFound)

		mockHasher.
			EXPECT().
			Hash("password123").
			Return("", expectedErr)

		register := Register{
			Repo:         mockRepo,
			Generator:    mockGenerator,
			HashPassword: mockHasher,
		}

		user := entities.User{
			Email:    "test@example.com",
			Name:     "John",
			Phone:    "08123456789",
			Password: "password123",
		}

		result, err := register.Execute(user)

		assert.ErrorIs(t, err, expectedErr)
		assert.Equal(t, entities.RegisteredUser{}, result)
	})

	t.Run("should return error when repository failed to create user", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockRepo := mocks.NewMockUserRepository(ctrl)
		mockGenerator := mocks.NewMockGeneratorId(ctrl)
		mockHasher := mocks.NewMockPasswordHasher(ctrl)

		expectedErr := errors.New("failed to create user")

		mockGenerator.
			EXPECT().
			Generator().
			Return("user-123")

		mockRepo.
			EXPECT().
			FindUserByEmail(gomock.Any()).
			Return(ErrNotFound)

		mockHasher.
			EXPECT().
			Hash("password123").
			Return("hashed-password", nil)

		mockRepo.
			EXPECT().
			UserRegister(gomock.Any()).
			Return(expectedErr)

		register := Register{
			Repo:         mockRepo,
			Generator:    mockGenerator,
			HashPassword: mockHasher,
		}

		user := entities.User{
			Email:    "test@example.com",
			Name:     "John",
			Phone:    "08123456789",
			Password: "password123",
		}

		result, err := register.Execute(user)

		assert.ErrorIs(t, err, expectedErr)
		assert.Equal(t, entities.RegisteredUser{}, result)
	})

	t.Run("create user success", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockRepo := mocks.NewMockUserRepository(ctrl)
		mockGenerator := mocks.NewMockGeneratorId(ctrl)
		mockHasher := mocks.NewMockPasswordHasher(ctrl)

		mockGenerator.
			EXPECT().
			Generator().
			Return("user-123")

		mockRepo.
			EXPECT().
			FindUserByEmail(gomock.Any()).
			Return(ErrNotFound)

		mockHasher.
			EXPECT().
			Hash("password123").
			Return("hashed-password", nil)

		mockRepo.
			EXPECT().
			UserRegister(gomock.Any()).
			DoAndReturn(func(user *entities.User) error {
				assert.Equal(t, "user-123", user.ID)
				assert.Equal(t, "test@example.com", user.Email)
				assert.Equal(t, "John", user.Name)
				assert.Equal(t, "08123456789", user.Phone)
				assert.Equal(t, "hashed-password", user.Password)

				return nil
			})

		register := Register{
			Repo:         mockRepo,
			Generator:    mockGenerator,
			HashPassword: mockHasher,
		}

		user := entities.User{
			Email:    "test@example.com",
			Name:     "John",
			Phone:    "08123456789",
			Password: "password123",
		}

		expected := entities.RegisteredUser{
			ID:   "user-123",
			Name: "John",
		}

		result, err := register.Execute(user)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}
