package ports

import "goWallet/internal/domain"

type AccountSrv interface {
	Register(*domain.Account) error
}

type AccountRepo interface {
	Register(*domain.Account) error
}
