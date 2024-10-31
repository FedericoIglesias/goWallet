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
	return s.Repo.Register(user)
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
