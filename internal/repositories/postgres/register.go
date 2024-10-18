package postgres

import (
	"goWallet/internal/domain"
)

func (r Repository) Register(user *domain.User) error {

	if err := r.Client.Create(user); err != nil {
		return err.Error
	}
	return nil
}
