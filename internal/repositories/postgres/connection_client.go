package postgres

import (
	"goWallet/internal/domain"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(DSN string) (DB *gorm.DB, err error) {

	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		return nil, err
	} else {
		log.Println("Connect up to DB")
	}

	DB.AutoMigrate(&domain.User{})

	return DB, nil
}
