package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/qwerty2265/go-chi-subscription-manager/internal/common/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("❌ Error loading .env file: %v", err)
	}
	log.Println("✅ .env file loaded successfully")
	serverPort := os.Getenv("SERVER_PORT")

	db.ConnectDB()

	log.Printf("🚀 Server is running on port %v", serverPort)
	if err := http.ListenAndServe(":"+serverPort, nil); err != nil {
		log.Fatalf("❌ The server failed to start: %v", err)
	}
}
