package auth

import (
	"errors"
	"os"
	"time"

	. "bitcoin-api-docker/api/src/customtypes"
	"bitcoin-api-docker/api/src/infrastructure/mysql"

	jwt "github.com/dgrijalva/jwt-go"
)

func generateJwtToken(account Account) (string, error) {
	jwtMethod := jwt.GetSigningMethod("HS256")
	claims := &jwt.StandardClaims{
		Subject:   account.UserID,
		ExpiresAt: (time.Now().Add(60 * time.Minute)).Unix(),
		// Issuer:    source,
	}

	token := jwt.NewWithClaims(jwtMethod, claims)
	signToken, err := token.SignedString([]byte(os.Getenv("PRIVATE_KEY")))

	if err != nil {
		return "", err
	}

	return signToken, nil
}

func Login(account Account) (map[string]string, error) {
	command := mysql.SelectCommand{
		TableName:     os.Getenv("USERS_DB"),
		Projection:    []string{"Password"},
		Join:          mysql.NONE,
		Conditions:    mysql.USER,
		ConditionData: []string{account.UserID},
	}
	result, err := mysql.Select(command)

	if err != nil {
		return nil, err
	}

	if result[0]["Password"] != account.Password {
		return nil, errors.New("User passaword doesn't match!")
	}

	jwtToken, err := generateJwtToken(account)

	if err != nil {
		return nil, err
	}

	return map[string]string{"Token": jwtToken}, nil
}
