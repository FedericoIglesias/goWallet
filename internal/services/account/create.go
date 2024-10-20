package servicesAccount

import (
	"goWallet/internal/domain"
)

func (s Services) Create(account *domain.Account) error {

	if err := s.Repo.Create(account); err != nil {
		return err
	}

	return nil
}
