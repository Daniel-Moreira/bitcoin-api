package auth

import (
	"errors"

	. "bitcoin-api/src/customtypes"
)

func Validate(user User) error {
	var result string
	if len(user.UserID) < 4 {
		result = "UserID must have at least 4 characters\n"
	}
	if len(user.Password) < 4 {
		result += "Password must have at least 4 characters\n"
	}

	if result == "" {
		return nil
	}

	return errors.New(result)
}
