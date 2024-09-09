package main

import (
	"notismart-service/internal/db"
	"notismart-service/internal/notification"
    "github.com/joho/godotenv"
	"time"
	"log"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	
	db.InitDB() // Initialize the database
	db.RunMigrations() // Run database migrations
	// Run a Goroutine to check for pending notifications every minute
	go func() {
		for {
			notification.DispatchPendingNotifications()
			time.Sleep(1 * time.Minute)
		}
	}()

	select {} // Keep the main function running
}
