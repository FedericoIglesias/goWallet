package models

type User struct {
	ID       int    `gorm: "unique, autoIncrement",json: "id"`
	Name     *string `json: "name"`
	LastName *string `json: "lastName"`
	Email    *string `gorm: "unique" ,json: "email"`
	Phone    *int    `json: "phone"`
	Password *string `json: "password"`
	Cvu      *uint   `gorm: "unique",json: "cvu"`
	Alias    *string `gorm: "unique",json: "alias"`
}
