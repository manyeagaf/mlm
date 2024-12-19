package database

import (
	"log"
	"mlm/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection
func ConnectDatabase() {
	var err error
	// Use SQLite as the database driver
	DB, err = gorm.Open(sqlite.Open("mlm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Migrate the schema
	err = DB.AutoMigrate(
		&models.User{},
		&models.Merchant{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database connected and migrated successfully!")
}
