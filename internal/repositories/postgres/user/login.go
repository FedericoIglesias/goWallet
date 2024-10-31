package postgresUser

import (
	"goWallet/internal/domain"
)

func (r Repository) GetUser(userLogin *domain.UserLogin) (domain.User , error) {

	var user domain.User

	err := r.Client.Where(&domain.User{Email: userLogin.Email}).First(&user)

	if err != nil {
		return user, err.Error
	}

	return user, nil
}
