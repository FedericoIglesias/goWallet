package postgresUser

import (
	"goWallet/internal/domain"
	// "github.com/pelletier/go-toml/query"
)

func (r Repository) Register(user *domain.User) (*domain.User, error) {
	query := `insert into users("name","last_name","email","phone","password","cvu","alias") values($1,$2,$3,$4,$5,$6,$7)`
	// `insert into users("name","last_name","email","phone","password","cvu","alias") values(?,?,?,?,?,?,?)`
	_, err := r.Client.Query(query, user.Name, user.LastName, user.Email, user.Phone, user.Password, user.Cvu, user.Alias)

	if err != nil {
		return &domain.User{}, err
	}
	// userLog := &domain.UserLogin{
	// 	Email:    user.Email,
	// 	Password: "",
	// }
	// user, err := r.GetUser(userLog)
	// if err != nil {
	// 	return &domain.User{}, err
	// }

	return user, nil
}
