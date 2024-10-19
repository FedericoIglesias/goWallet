package domain

type User struct {
	ID       uint    `gorm:"unique, autoIncrement" json:"id"`
	Name     string  `gorm:"not null" json:"name"`
	LastName string  `gorm:"not null" json:"last_name"`
	Email    string  `gorm:"not null ;unique"  json:"email"`
	Phone    int     `gorm:"not null" json:"phone"`
	Password string  `gorm:"not null" json:"password"`
	Cvu      string  `gorm:"not null ;unique" json:"cvu"`
	Alias    string  `gorm:"not null ;unique" json:"alias"`
	Account  Account `gorm:"not null; foreignKey:Email"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
