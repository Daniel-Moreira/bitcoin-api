package auth

import (
	"errors"
	"fmt"
	"os"

	. "bitcoin-api-docker/api/src/customtypes"
	"bitcoin-api-docker/api/src/infrastructure/mysql"

	jwt "github.com/dgrijalva/jwt-go"
)

func generateJwtToken(account Account, source string) (string, error) {
	jwtMethod := jwt.GetSigningMethod("HS256")
	claims := &jwt.StandardClaims{
		Subject:   account.UserID,
		ExpiresAt: 15000,
		Issuer:    source,
	}

	token := jwt.NewWithClaims(jwtMethod, claims)
	signToken, err := token.SignedString([]byte(os.Getenv("PRIVATE_KEY")))

	if err != nil {
		return "", err
	}

	return signToken, nil
}

func Login(account Account, source string) (string, error) {
	if source == "" {
		return "", errors.New("Source is required")
	}

	command := mysql.SelectCommand{
		TableName:     os.Getenv("USERS_DB"),
		Projection:    []string{"password"},
		Join:          mysql.NONE,
		Conditions:    mysql.USER,
		ConditionData: []string{account.UserID},
	}
	result, err := mysql.Select(command)

	if err != nil {
		return "", err
	}

	fmt.Println(result[0]["password"])
	if result[0]["password"] != account.Password {
		return "", errors.New("User passaword doesn't match!")
	}

	jwtToken, err := generateJwtToken(account, source)

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
