package services

import (
	"goWallet/internal/db"
	"goWallet/internal/models"
)

func Register(user *models.User) error {
	return db.DB.Create(&user).Error
}
