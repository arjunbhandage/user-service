package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	return db
}
