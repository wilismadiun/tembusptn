package driver

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"subscription-service/commons/database"
	"subscription-service/src/Domains/entities"
	"subscription-service/src/Infrastructures/security"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Println("Peringatan: Gagal memuat .env dari root, mencoba path alternatif...")
	}

	database.ConnectDatabase()

	gin.SetMode(gin.TestMode)

	code := m.Run()

	os.Exit(code)
}

func TestAddSubscription(t *testing.T) {
	t.Run("should return 401 when authorization header is empty", func(t *testing.T) {
		router := gin.New()
		Router(router, database.DB)

		body := []byte(`{
			"name": "Gold",
			"price": 50000
		}`)

		req := httptest.NewRequest(
			http.MethodPost,
			"/subscriptions",
			bytes.NewBuffer(body),
		)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	})

	t.Run("should return 403 when role is not admin", func(t *testing.T) {
		router := gin.New()
		Router(router, database.DB)

		tokenGenerator := security.AuthenticationTokenJWT{}

		token, err := tokenGenerator.GenerateToken("user-123", "student")
		require.NoError(t, err)

		body := []byte(`{
			"name": "Gold",
			"price": 50000
		}`)

		req := httptest.NewRequest(
			http.MethodPost,
			"/subscriptions",
			bytes.NewBuffer(body),
		)

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusForbidden, recorder.Code)
	})

	t.Run("should return 400 when name is empty", func(t *testing.T) {
		router := gin.New()
		Router(router, database.DB)

		tokenGenerator := security.AuthenticationTokenJWT{}

		token, err := tokenGenerator.GenerateToken("user-123", "admin")
		require.NoError(t, err)

		body := []byte(`{
			"name": "",
			"price": 50000
		}`)

		req := httptest.NewRequest(
			http.MethodPost,
			"/subscriptions",
			bytes.NewBuffer(body),
		)

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("should return 400 when price is empty", func(t *testing.T) {
		router := gin.New()
		Router(router, database.DB)

		tokenGenerator := security.AuthenticationTokenJWT{}

		token, err := tokenGenerator.GenerateToken("user-123", "admin")
		require.NoError(t, err)

		body := []byte(`{
			"name": "Gold",
			"price": 0
		}`)

		req := httptest.NewRequest(
			http.MethodPost,
			"/subscriptions",
			bytes.NewBuffer(body),
		)

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("should return 400 when subscription name already exists", func(t *testing.T) {
		router := gin.New()
		Router(router, database.DB)

		dummySubscription := entities.Subscriptions{
			ID:    "subscription-existing-123",
			Name:  "Gold Existing",
			Price: 50000,
		}

		err := database.DB.Create(&dummySubscription).Error
		require.NoError(t, err)

		tokenGenerator := security.AuthenticationTokenJWT{}

		token, err := tokenGenerator.GenerateToken("user-123", "admin")
		require.NoError(t, err)

		body := []byte(`{
			"name": "Gold Existing",
			"price": 100000
		}`)

		req := httptest.NewRequest(
			http.MethodPost,
			"/subscriptions",
			bytes.NewBuffer(body),
		)

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		database.DB.
			Where("id = ?", dummySubscription.ID).
			Delete(&entities.Subscriptions{})
	})

	t.Run("should return 201 when subscription successfully created", func(t *testing.T) {
		router := gin.New()
		Router(router, database.DB)

		tokenGenerator := security.AuthenticationTokenJWT{}

		token, err := tokenGenerator.GenerateToken("user-123", "admin")
		require.NoError(t, err)

		body := []byte(`{
			"name": "Diamond Integration Test",
			"price": 150000
		}`)

		req := httptest.NewRequest(
			http.MethodPost,
			"/subscriptions",
			bytes.NewBuffer(body),
		)

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusCreated, recorder.Code)

		var result entities.Subscriptions

		err = database.DB.
			Where("name = ?", "Diamond Integration Test").
			First(&result).
			Error

		assert.NoError(t, err)
		assert.Equal(t, "Diamond Integration Test", result.Name)
		assert.Equal(t, int64(150000), result.Price)

		database.DB.
			Where("name = ?", "Diamond Integration Test").
			Delete(&entities.Subscriptions{})
	})
}

func TestGetAllSubscriptions(t *testing.T) {

	t.Run("should return 200 when data is empty", func(t *testing.T) {
		database.DB.Exec("DELETE FROM subscriptions")

		router := gin.New()
		Router(router, database.DB)

		req := httptest.NewRequest(
			http.MethodGet,
			"/subscriptions",
			nil,
		)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.JSONEq(t, `{
			"message": "Data not found",
			"data": []
		}`, recorder.Body.String())
	})

	t.Run("should return 200 when data exists", func(t *testing.T) {
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
		require.NoError(t, err)

		router := gin.New()
		Router(router, database.DB)

		req := httptest.NewRequest(
			http.MethodGet,
			"/subscriptions",
			nil,
		)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var response map[string]interface{}

		err = json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, recorder.Code)

		assert.JSONEq(t, `{
		"message": "Success",
		"data": [
			{
				"id": "subscription-1",
				"name": "Gold",
				"price": 50000
			},
			{
				"id": "subscription-2",
				"name": "Diamond",
				"price": 150000
			}
		]
	}`, recorder.Body.String())

		database.DB.Exec("DELETE FROM subscriptions")
	})
}

func TestAddUserSubScripitions(t *testing.T) {

	t.Run("should return 401 when jwt is not provided", func(t *testing.T) {
		router := gin.New()
		Router(router, database.DB)

		body := []byte(`{
			"subscription_id": "subscription-123",
			"duration_in_days": 30
		}`)

		req := httptest.NewRequest(
			http.MethodPost,
			"/user-subscriptions",
			bytes.NewBuffer(body),
		)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	})

	t.Run("should return 400 when subscription id is empty", func(t *testing.T) {
		router := gin.New()
		Router(router, database.DB)

		tokenGenerator := security.AuthenticationTokenJWT{}

		token, err := tokenGenerator.GenerateToken("user-123", "admin")
		require.NoError(t, err)

		body := []byte(`{
			"subscription_id": "",
			"duration_in_days": 30
		}`)

		req := httptest.NewRequest(
			http.MethodPost,
			"/user-subscriptions",
			bytes.NewBuffer(body),
		)

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("should return 400 when duration is less than or equal to 0", func(t *testing.T) {
		router := gin.New()
		Router(router, database.DB)

		tokenGenerator := security.AuthenticationTokenJWT{}

		token, err := tokenGenerator.GenerateToken("user-123", "admin")
		require.NoError(t, err)

		body := []byte(`{
			"subscription_id": "subscription-123",
			"duration_in_days": 0
		}`)

		req := httptest.NewRequest(
			http.MethodPost,
			"/user-subscriptions",
			bytes.NewBuffer(body),
		)

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("should return 201 when user subscription successfully created", func(t *testing.T) {
		router := gin.New()
		Router(router, database.DB)

		// Data subscription
		subscription := entities.Subscriptions{
			ID:    "subscription-123",
			Name:  "Gold Integration Test",
			Price: 50000,
		}

		err := database.DB.
			Where("id = ?", subscription.ID).
			FirstOrCreate(&subscription).
			Error

		require.NoError(t, err)

		defer database.DB.
			Where("id = ?", subscription.ID).
			Delete(&entities.Subscriptions{})

		// Generate JWT
		tokenGenerator := security.AuthenticationTokenJWT{}

		token, err := tokenGenerator.GenerateToken("user-123", "admin")
		require.NoError(t, err)

		body := []byte(`{
		"subscription_id": "subscription-123",
		"duration_in_days": 1
	}`)

		req := httptest.NewRequest(
			http.MethodPost,
			"/user-subscriptions",
			bytes.NewBuffer(body),
		)

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusCreated, recorder.Code)

		// Cek data yang masuk ke database
		var result entities.UserSubscriptions

		err = database.DB.
			Where(
				"user_id = ? AND subscription_id = ?",
				"user-123",
				"subscription-123",
			).
			First(&result).
			Error

		require.NoError(t, err)

		assert.Equal(t, "user-123", result.UserId)
		assert.Equal(t, "subscription-123", result.SubscriptionId)
		assert.Equal(t, entities.StatusActive, result.Status)

		assert.False(t, result.StartAt.IsZero())
		assert.False(t, result.EndAt.IsZero())

		// duration_in_days = 30
		assert.Equal(
			t,
			result.StartAt.AddDate(0, 0, 30),
			result.EndAt,
		)

		// Hapus data user subscription
		database.DB.
			Where("id = ?", result.ID).
			Delete(&entities.UserSubscriptions{})
	})
}
