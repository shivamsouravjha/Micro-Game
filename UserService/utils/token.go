package utils

import (
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("mySigningKey"))
	return tokenString, err
}
