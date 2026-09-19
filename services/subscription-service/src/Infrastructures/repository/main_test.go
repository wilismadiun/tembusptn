package repository

import (
	"log"
	"os"
	"subscription-service/commons/database"
	"testing"

	"github.com/joho/godotenv"
)

var subscriptionRepo *SubscriptionRepository
var userSubsRepo *UserSubscriptionRepository

func TestMain(m *testing.M) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Println("Peringatan: Gagal memuat .env dari root")
	}

	database.ConnectDatabase()

	subscriptionRepo = &SubscriptionRepository{
		DB: database.DB,
	}

	userSubsRepo = &UserSubscriptionRepository{
		DB: database.DB,
	}

	code := m.Run()

	os.Exit(code)
}
