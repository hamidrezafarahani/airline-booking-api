package main

import (
	"log"

	"github.com/hamidrezafarahani/airline-booking-api/internal/config"
	"github.com/hamidrezafarahani/airline-booking-api/internal/infrastructure/persistence/postgres"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("warning: .env file not found")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.New(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	if err := postgres.Migrate(db); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	if err := postgres.Seed(db); err != nil {
		log.Fatalf("seed failed: %v", err)
	}

	log.Println("database initialized successfully")
}
