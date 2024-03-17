package database

import (
	"log"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB, dst ...interface{}) {
	log.Println("Running migrations....")
	db.AutoMigrate(dst...)
}
