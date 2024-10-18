package services

import (
	"fmt"
	"goWallet/internal/domain"
)

func (s Services) Login(userLogin *domain.UserLogin) error {

	user, err := s.Repo.GetUser(userLogin)

	if err != nil {
		return err
	}

	if userLogin.Password != user.Password {
		return fmt.Errorf("incorrect password", nil) // ("incorrect password")
	}

	return nil
}
