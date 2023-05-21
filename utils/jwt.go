package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY = "SECRET_T0KEN"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	webtoken, err := token.SignedString([]byte(SECRET_KEY))

	if err != nil {
		return "", err
	}

	return webtoken, nil

}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil

}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	VerifyToken()
}
