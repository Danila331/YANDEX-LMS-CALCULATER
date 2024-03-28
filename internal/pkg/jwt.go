package pkg

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = "mami22"

func MakeJwt(login string) (string, error) {
	// Создаем структуру токена с именем пользователя в качестве пользовательского клейма
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["login"] = login
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Токен действителен в течение 24 часов

	// Получаем строковое представление токена
	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidJwt(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return false, err
	}

	// Проверка, является ли токен валидным и не истек
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		return time.Now().After(expirationTime), nil
	}
	return true, nil
}
