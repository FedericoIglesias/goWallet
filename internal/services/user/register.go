package servicesUser

import (
	"fmt"
	"goWallet/internal/domain"
	"os"
	"strings"

	"math/rand"
)

func (s Services) Register(user *domain.User) error {
	user.Cvu = CreateCVU()
	user.Alias = CreateAlias()
	user, err := s.Repo.Register(user)
	// Create Account

	s.RepoAccount.Create(&domain.Account{Cash: 0.0, UserID: user.ID})
	if err != nil {
		return err
	}

	return nil
}

func CreateCVU() string {
	formatted := fmt.Sprintf("%019d", rand.Uint64())

	return ("996" + formatted)
}

func CreateAlias() string {
	file, err := os.ReadFile("./internal/services/alia.csv")
	if err != nil {
		return err.Error()
	}

	listWord := strings.Split(string(file), "\n")

	alias := listWord[rand.Intn(100)]

	return alias

}
