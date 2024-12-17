package postgres

import (
	"database/sql"
	"goWallet/internal/domain"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(DSN string) (DB *sql.DB, err error) {

	MG, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		return nil, err
	} else {
		log.Println("Connect up to DB")
	}
	MG.AutoMigrate(&domain.User{}, &domain.Account{})

	DB, err = sql.Open("postgres", DSN)
	if err != nil {
		return nil, err
	} else {
		log.Println("Connect up to DB")
	}
	return DB, nil
}
