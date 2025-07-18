package app

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/qwerty2265/go-chi-subscription-manager/internal/common/db"
	"github.com/qwerty2265/go-chi-subscription-manager/internal/subscription"
)

func InitializeApp() chi.Router {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("❌ Error loading .env file: %v", err)
	}
	log.Println("✅ .env file loaded successfully")

	database := db.ConnectDB()
	db.Migrate(database)

	subRepo := subscription.NewSubscriptionRepository(database)
	subService := subscription.NewSubscriptionService(subRepo)
	subHandler := subscription.NewSubscriptionHandler(subService)

	router := NewRouter(subHandler)

	log.Println("✅ Application initialized successfully")
	return router
}
