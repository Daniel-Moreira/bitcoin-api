package auth

import (
	"errors"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

func authenticateJwtToken(token string) (string, error) {
	claims := &jwt.StandardClaims{}

	tk, err := jwt.ParseWithClaims(token, claims, func(tk *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("PRIVATE_KEY")), nil
	})

	if err != nil {
		return "", err
	}

	if !tk.Valid {
		return "", errors.New("Unauthorized access!")
	}

	return claims.Subject, nil
}

func Authenticate(token string) (string, error) {
	if token == "" {
		return "", errors.New("Missing authentication token")
	}

	userId, err := authenticateJwtToken(token)

	if err != nil {
		return "", err
	}

	return userId, nil
}
