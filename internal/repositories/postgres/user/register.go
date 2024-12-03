package postgresUser

import (
	"goWallet/internal/domain"
)

func (r Repository) Register(user *domain.User) (*domain.User, error) {

	if err := r.Client.Create(user); err != nil {
		return &domain.User{}, err.Error
	}
	userLog := &domain.UserLogin{
		Email:    user.Email,
		Password: "",
	}
	user, err := r.GetUser(userLog)
	if err != nil {
		return &domain.User{}, err
	}

	return user, nil
}
