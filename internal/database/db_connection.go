package database

import (
	"disbursement-service/internal/config"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDatabaseConnection(config *config.Config) *gorm.DB {
	dbName := fmt.Sprintf("%s.db", config.Database.Name)
	connection, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database \n", err.Error())
		os.Exit(1)
	}

	return connection
}
