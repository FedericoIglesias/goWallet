package domain

type Account struct {
	ID    uint    `gorm:"autoincrement,unique" json:"id"`
	Cash  float64 `gorm:"not null" json:"cash"`
	Email string  `json:"email"`
}
