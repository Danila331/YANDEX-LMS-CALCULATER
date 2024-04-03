package pkg

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type MyCustomClaims struct {
	UserLogin string `json:"userlogin"`
	jwt.StandardClaims
}

var secretKey = []byte("mami")

func GenerateJWT(userLogin string) (string, error) {
	claims := MyCustomClaims{
		userLogin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Пример: токен действителен 24 часа                  // Замените "your-issuer" на ваше имя или идентификатор приложения
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserLoginFromJWT(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid JWT token")
	}

	return claims.UserLogin, nil
}
