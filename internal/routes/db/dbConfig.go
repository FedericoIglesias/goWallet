package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := "host=localhost user=postgres password=admin dbname=finance port=5432 sslmode=disable TimeZone=UTC"

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error to connect to database")
	}

	DB.AutoMigrate()

}
