package database

import (
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(url string) {
	// Initialize the migrate instance
	m, err := migrate.New(
		"file://db/migrations",
		url,
	)
	if err != nil {
		log.Fatalf("Failed to initialize migration: %v", err)
	}

	// Run migrations up
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migrations applied successfully.")
}
