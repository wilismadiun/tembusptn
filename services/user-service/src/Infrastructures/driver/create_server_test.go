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
	"github.com/wilismadiun/tembusptn/services/user-service/commons/database"
	"github.com/wilismadiun/tembusptn/services/user-service/src/Domains/entities"
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

func databaseHelper() {
	database.DB.Exec("DELETE FROM users")

	user := entities.User{
		ID:       "user-123",
		RoleId:   2,
		Email:    "john@gmail.com",
		Phone:    "000000000000",
		Name:     "John",
		Password: "12345678",
	}

	err := database.DB.Create(&user).Error
	if err != nil {
		log.Printf("Error: %s", err)
	}
}

func Test_register(t *testing.T) {
	router := gin.New()

	Router(router, database.DB)

	t.Run("should response 400 when email already exists", func(t *testing.T) {
		database.DB.Exec("DELETE FROM users")

		databaseHelper()

		body := []byte(`{
			"roleId": 2,
			"email": "john@gmail.com",
			"name": "Jaya",
			"phone": "081234567890",
			"password": "12345678"
		}`)

		req := httptest.NewRequest(
			http.MethodPost,
			"/register",
			bytes.NewBuffer(body),
		)

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		database.DB.Exec("DELETE FROM users")
	})

	t.Run("should response 400 when password has less than 8 characters", func(t *testing.T) {
		database.DB.Exec("DELETE FROM users")

		body := []byte(`{
			"roleId": 2,
			"email": "jaya@gmail.com",
			"name": "Jaya",
			"phone": "081234567890",
			"password": "1234567"
		}`)

		req := httptest.NewRequest(
			http.MethodPost,
			"/register",
			bytes.NewBuffer(body),
		)

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		database.DB.Exec("DELETE FROM users")
	})

	t.Run("should response 400 when role id is empty", func(t *testing.T) {
		database.DB.Exec("DELETE FROM users")

		body := []byte(`{
			"roleId": 0,
			"email": "jaya@gmail.com",
			"name": "Jaya",
			"phone": "081234567890",
			"password": "12345678"
		}`)

		req := httptest.NewRequest(
			http.MethodPost,
			"/register",
			bytes.NewBuffer(body),
		)

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		database.DB.Exec("DELETE FROM users")
	})

	t.Run("registration successful", func(t *testing.T) {
		database.DB.Exec("DELETE FROM users")

		body := []byte(`{
			"roleId": 2,
			"email": "jaya@gmail.com",
			"name": "Jaya",
			"phone": "081234567890",
			"password": "12345678"
		}`)

		req := httptest.NewRequest(
			http.MethodPost,
			"/register",
			bytes.NewBuffer(body),
		)

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		// Check database
		var user entities.User

		err := database.DB.
			Where("email = ?", "jaya@gmail.com").
			First(&user).
			Error

		assert.NoError(t, err)

		assert.Equal(t, "Jaya", user.Name)
		assert.Equal(t, "jaya@gmail.com", user.Email)
		assert.Equal(t, "081234567890", user.Phone)
		assert.Equal(t, 2, user.RoleId)

		// Password should be hashed
		assert.NotEqual(t, "12345678", user.Password)

		// Password should not be empty
		assert.NotEmpty(t, user.Password)

		// Check response
		var response map[string]interface{}

		err = json.Unmarshal(
			w.Body.Bytes(),
			&response,
		)

		assert.NoError(t, err)
		assert.NotEmpty(t, response)

		database.DB.Exec("DELETE FROM users")
	})
}

func Test_login(t *testing.T) {
	router := gin.New()

	Router(router, database.DB)

	t.Run("should response 400 when email is empty", func(t *testing.T) {
		database.DB.Exec("DELETE FROM users")

		body := []byte(`{
			"email": "",
			"password": "12345678"
		}`)

		req := httptest.NewRequest(
			http.MethodPost,
			"/login",
			bytes.NewBuffer(body),
		)

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should response 400 when email is not registered", func(t *testing.T) {
		database.DB.Exec("DELETE FROM users")

		body := []byte(`{
			"email": "notregistered@gmail.com",
			"password": "12345678"
		}`)

		req := httptest.NewRequest(
			http.MethodPost,
			"/login",
			bytes.NewBuffer(body),
		)

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should response 400 when password is incorrect", func(t *testing.T) {
		database.DB.Exec("DELETE FROM users")

		// Register user first
		registerBody := []byte(`{
			"roleId": 2,
			"email": "login@gmail.com",
			"name": "Jaya",
			"phone": "081234567890",
			"password": "12345678"
		}`)

		registerReq := httptest.NewRequest(
			http.MethodPost,
			"/register",
			bytes.NewBuffer(registerBody),
		)

		registerReq.Header.Set(
			"Content-Type",
			"application/json",
		)

		registerW := httptest.NewRecorder()

		router.ServeHTTP(registerW, registerReq)

		assert.Equal(t, http.StatusCreated, registerW.Code)

		// Login with wrong password
		loginBody := []byte(`{
			"email": "login@gmail.com",
			"password": "wrongpassword"
		}`)

		loginReq := httptest.NewRequest(
			http.MethodPost,
			"/login",
			bytes.NewBuffer(loginBody),
		)

		loginReq.Header.Set(
			"Content-Type",
			"application/json",
		)

		loginW := httptest.NewRecorder()

		router.ServeHTTP(loginW, loginReq)

		assert.Equal(t, http.StatusBadRequest, loginW.Code)

		database.DB.Exec("DELETE FROM users")
	})

	t.Run("login successful", func(t *testing.T) {
		database.DB.Exec("DELETE FROM users")

		// Register user first
		registerBody := []byte(`{
			"roleId": 2,
			"email": "success@gmail.com",
			"name": "Jaya",
			"phone": "081234567890",
			"password": "12345678"
		}`)

		registerReq := httptest.NewRequest(
			http.MethodPost,
			"/register",
			bytes.NewBuffer(registerBody),
		)

		registerReq.Header.Set(
			"Content-Type",
			"application/json",
		)

		registerW := httptest.NewRecorder()

		router.ServeHTTP(registerW, registerReq)

		assert.Equal(t, http.StatusCreated, registerW.Code)

		// Login
		loginBody := []byte(`{
			"email": "success@gmail.com",
			"password": "12345678"
		}`)

		loginReq := httptest.NewRequest(
			http.MethodPost,
			"/login",
			bytes.NewBuffer(loginBody),
		)

		loginReq.Header.Set(
			"Content-Type",
			"application/json",
		)

		loginW := httptest.NewRecorder()

		router.ServeHTTP(loginW, loginReq)

		assert.Equal(t, http.StatusOK, loginW.Code)

		// Check response
		var response map[string]interface{}

		err := json.Unmarshal(
			loginW.Body.Bytes(),
			&response,
		)

		assert.NoError(t, err)
		assert.NotEmpty(t, response)

		database.DB.Exec("DELETE FROM users")
	})
}
