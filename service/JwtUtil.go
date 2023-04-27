package service

import (
	"fmt"
	"log"
	"portofolio/config"
	"portofolio/model"
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

func ValidateJWT(tokenString string) (model.UserInfo, error) {
	secretKey := config.Configuration().JwtSecretKey

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing Method Invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Signing Method Invalid")
		}

		return []byte(secretKey), nil
	})

	myClaim, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		log.Println(myClaim["username"])
	}

	log.Println(token.Claims)

	if err != nil {
		return model.UserInfo{}, err
	}

	return model.UserInfo{
		Username: myClaim["username"].(string),
	}, nil
}
