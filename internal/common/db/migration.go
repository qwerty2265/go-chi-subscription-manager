package db

import (
	"log"

	"github.com/qwerty2265/go-chi-subscription-manager/internal/subscription"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&subscription.Subscription{},
	)

	if err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}

	log.Println("🚚 Migrations has been successfully applied")
}
