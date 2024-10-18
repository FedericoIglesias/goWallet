package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY = os.Getenv("SECRET_KEY")

var jwtKey = []byte(SECRET_KEY)

func GenerateToken(email *string, id *uint) (string, error) {

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 1)),
		Subject:   *email,
		ID:        string(rune(*id)),
	}

	preToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := preToken.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return token, nil

}
