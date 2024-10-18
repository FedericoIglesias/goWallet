package ports

import "goWallet/internal/domain"

type UserSrv interface {
	Register(user *domain.User) error
}

type UserRepo interface {
	Register(user *domain.User) error
}