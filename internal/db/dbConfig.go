package db

import (
	"goWallet/internal/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	var err error

	DSN := os.Getenv("DSN")

	dsn := DSN

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error to connect to database")
	}else {
		log.Println("Connect up to DB")
	}

	DB.AutoMigrate(&models.User{})

}
