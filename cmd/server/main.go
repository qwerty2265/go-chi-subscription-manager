package main

import (
	"log"
	"net/http"
	"os"

	"github.com/qwerty2265/go-chi-subscription-manager/app"
)

func main() {
	router := app.InitializeApp()
	serverPort := os.Getenv("SERVER_PORT")

	log.Printf("🚀 Server is running on port %v", serverPort)
	if err := http.ListenAndServe(":"+serverPort, router); err != nil {
		log.Fatalf("❌ The server failed to start: %v", err)
	}
}
