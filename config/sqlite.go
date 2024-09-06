package config

import (
	"os"

	"github.com/fabiopasilva1/gooportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger(("sqlite"))
	dbPath := "./db/main.db"
	// Check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file does not exist. Creating it...")
		err = os.MkdirAll("./db", os.ModePerm)
		// Create the database file

		if err != nil {
			logger.Errorf("Error creating database file: %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Create DB sqlite
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error connecting to database: %v", err)
		return nil, err
	}

	// Migrate DB Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error migrating database: %v", err)
		return nil, err
	}
	return db, nil
}
