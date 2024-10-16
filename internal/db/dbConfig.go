package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	DSN := os.Getenv("DSN")

	var err error

	dsn := DSN

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error to connect to database")
	}

	DB.AutoMigrate()

}
