package services

import (
	"fmt"
	"goWallet/internal/domain"
	"goWallet/internal/utils/jwt"
)

func (s Services) Login(userLogin *domain.UserLogin) (string, error) {

	user, err := s.Repo.GetUser(userLogin)

	if err != nil {
		return "", err
	}

	if userLogin.Password != user.Password {
		return "", fmt.Errorf("incorrect password", nil) // ("incorrect password")
	}

	return jwt.GenerateToken(&user.Email, &user.ID)
}
