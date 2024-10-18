package ports

import "goWallet/internal/domain"

type UserSrv interface {
	Register(user *domain.User) error
	Login(user *domain.UserLogin) (string, error)
}

type UserRepo interface {
	Register(user *domain.User) error
	GetUser(user *domain.UserLogin) (domain.User, error)
}
