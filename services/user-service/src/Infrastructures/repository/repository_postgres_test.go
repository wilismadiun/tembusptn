package repository

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/wilismadiun/tembusptn/services/user-service/commons/database"
	"github.com/wilismadiun/tembusptn/services/user-service/src/Applications/usecase"
	"github.com/wilismadiun/tembusptn/services/user-service/src/Domains/entities"
)

var repo *UserRepository

func TestMain(m *testing.M) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Println("Peringatan: Gagal memuat .env dari root, mencoba path alternatif...")
	}

	database.ConnectDatabase()

	repo = &UserRepository{
		DB: database.DB,
	}

	code := m.Run()

	os.Exit(code)
}

func TestFindUserByEmail_NotFound(T *testing.T) {
	nonExistentUser := entities.User{
		Email: "email_yang_pasti_tidak_ada@test.com",
	}

	_, err := repo.FindUserByEmail(nonExistentUser.Email)

	// Validasi bahwa error yang dikembalikan adalah usecase.ErrNotFound
	assert.Error(T, err)
	assert.True(T, errors.Is(err, usecase.ErrNotFound))
}

func TestFindUserByEmail_Found(t *testing.T) {
	dummyUser := entities.User{
		ID:       "uuid-123",
		Name:     "Test User",
		Email:    "test_found@test.com",
		Password: "hashedpassword",
	}

	err := database.DB.Create(&dummyUser).Error
	assert.NoError(t, err)

	searchUser := entities.User{
		Email: "test_found@test.com",
	}

	_, err = repo.FindUserByEmail(searchUser.Email)

	assert.NoError(t, err)

	database.DB.
		Where("email = ?", "test_found@test.com").
		Delete(&entities.User{})
}

func Test_UserRegister(t *testing.T) {

	user := entities.User{
		ID:       "user-register-123",
		Email:    "jaya@gmail.com",
		Name:     "Jaya",
		Phone:    "081234567890",
		Password: "hashed-password",
	}

	err := repo.UserRegister(&user)

	assert.NoError(t, err)

	var result entities.User

	err = database.DB.
		Where("email = ?", user.Email).
		First(&result).
		Error

	assert.NoError(t, err)

	assert.Equal(t, user.ID, result.ID)
	assert.Equal(t, user.Email, result.Email)
	assert.Equal(t, user.Name, result.Name)
	assert.Equal(t, user.Phone, result.Phone)
	assert.Equal(t, user.Password, result.Password)

	t.Logf("User berhasil disimpan: %+v", result)

	database.DB.Exec("DELETE FROM users")
}
