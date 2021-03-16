package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var mySigningKey = []byte("Bambi")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "bambi"
	claims["exp"] = time.Now().Add(time.Minute * 2000).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("error: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func main() {
	token, _ := GenerateJWT()
	print(token)
}
