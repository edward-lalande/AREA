package utils

import (
	"github.com/golang-jwt/jwt"
)

func CreateToken(username string) (string, error) {
	var secretKey string = GetEnvKey("SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
