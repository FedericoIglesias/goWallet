package postgresAccount

import (
	"fmt"
	"goWallet/internal/domain"
)

func (r Repository) Get(account *domain.Account) (*domain.Account, error) {
	err := r.Client.First(&account).Error

	if err != nil {
		return account, fmt.Errorf("error 500: %w", err)
	}
	return account, nil
}
