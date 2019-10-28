package auth

import (
	"errors"

	. "bitcoin-api-docker/api/src/customtypes"
)

func Validate(account Account) error {
	var result string
	if len(account.UserID) < 4 {
		result = "UserID must have at least 4 characters\n"
	}
	if len(account.Password) < 4 {
		result += "Password must have at least 4 characters\n"
	}

	if result == "" {
		return nil
	}

	return errors.New(result)
}
