package auth

import (
	"os"

	. "bitcoin-api/src/customtypes"
	. "bitcoin-api/src/infrastructure/aws/dynamo"
)

func Sign(user User) (map[string]string, error) {
	err := Validate(user)

	if err != nil {
		return nil, err
	}

	err = Put(os.Getenv("REGISTER_USERS"), user)

	if err != nil {
		return nil, err
	}

	return map[string]string{"Message": "Account Created!"}, nil
}
