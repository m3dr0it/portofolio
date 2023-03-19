package service

import (
	"fmt"
	"portofolio/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(username string) (string, error) {
	expired := config.Configuration().Server.JwtExpired
	secretKey := config.Configuration().Server.JwtSecretKey
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
