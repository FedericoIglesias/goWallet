package ports

import "goWallet/internal/domain"

type User interface {
	Register(user *domain.User) error
}