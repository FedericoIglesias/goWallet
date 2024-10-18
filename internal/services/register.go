package services

import (
	"fmt"
	"goWallet/internal/domain"

	"math/rand"
)

func (s Services) Register(user *domain.User) error {
	user.Cvu = CreateCVU()
	return s.Repo.Register(user)
}

func CreateCVU() string {
	formatted := fmt.Sprintf("%019d", rand.Uint64())

	return ("996" + formatted)
}
