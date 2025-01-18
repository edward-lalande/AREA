package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id string) (string, error) {
	var secretKey string = GetEnvKey("SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id": id,
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) string {
	var secretKey string = GetEnvKey("SECRET_KEY")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	fmt.Println("tokenString: ", tokenString)
	fmt.Println("token parse: ", token)
	if err != nil {
		fmt.Println("error:", err.Error())
		return ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("claims: ", claims)
		return claims["id"].(string)
	}

	return ""
}
