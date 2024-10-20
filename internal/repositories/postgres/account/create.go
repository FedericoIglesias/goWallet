package postgresAccount

import "goWallet/internal/domain"

func (r Repository) Create(account *domain.Account) error {

	if err := r.Client.Create(account); err != nil {
		return err.Error
	}
	return nil
}
