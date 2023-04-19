package service

import (
	"fmt"
	"log"
	"portofolio/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(username string) (string, error) {
	expired := config.Configuration().JwtExpired
	secretKey := config.Configuration().JwtSecretKey
	port := config.Configuration().Server.Port
	log.Println(port)
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Duration(expired) * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))

	if err != nil {
		fmt.Println("Error signing token:", err)
		return "", err
	}

	fmt.Println("JWT:", signedToken)

	return signedToken, nil
}

func ValidateJWT(tokenString string) (bool, error) {
	secretKey := config.Configuration().JwtSecretKey

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing Method Invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Signing Method Invalid")
		}

		return []byte(secretKey), nil
	})

	log.Println(token)

	if err != nil {
		return false, err
	}

	return true, nil
}
