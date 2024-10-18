package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(tokenString string) (*jwt.RegisteredClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)

	if !ok {
		return nil, fmt.Errorf("problem with claims in jwt")
	}

	return claims, nil
}
