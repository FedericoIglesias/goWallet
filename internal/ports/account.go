package ports

import "goWallet/internal/domain"

type AccountSrv interface {
	Create(*domain.Account) error
	Get(*domain.Account) (*domain.Account, error)
}

type AccountRepo interface {
	Create(*domain.Account) error
	Get(*domain.Account) (*domain.Account, error)
}
