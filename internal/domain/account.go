package domain

type Account struct {
	ID     uint    `gorm:"autoincrement;unique" json:"id"`
	Cash   float64 `gorm:"not null" json:"cash"`
	UserID uint    `gorm:"not null; unique" json:"user_id"`
}
