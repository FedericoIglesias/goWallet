package postgres

import (
	"gorm.io/gorm"
)

type Repository struct {
	Client *gorm.DB
}
