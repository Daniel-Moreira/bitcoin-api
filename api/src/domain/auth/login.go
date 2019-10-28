package auth

import (
  "errors"
	"os"

	"bitcoin-api/src/infrastructure/mysql"
	. "bitcoin-api/src/customtypes"

	jwt "github.com/dgrijalva/jwt-go"
)

func generateJwtToken(user User, source string) (string, error) {
	jwtMethod := jwt.GetSigningMethod("HS256")
	claims := &jwt.StandardClaims{
		Subject:   user.UserID,
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

func Login(user User, source string) (string, error) {
	if source == "" {
		return "", errors.New("Source is required")
	}

  command := mysql.SelectCommand{
    TableName: os.Getenv("REGISTER_USERS"),
    Projection: ["password"],
    Join: mysql.Join.None,
    Conditions: mysql.Conditions.User,
    ConditionData [user.UserID],
  }
	result, err := mysql.Select(command)

	if err != nil {
		return "", err
	}

  fmt.Println(result["password"])
	if result["password"] != user.Password {
		return "", errors.New("User passaword doesn't match!")
	}

	jwtToken, err := generateJwtToken(user, source)

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
