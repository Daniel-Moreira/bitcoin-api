package auth

import (
	"errors"
	"time"

	. "bitcoin-api-docker/api/src/customtypes"
)

func Validate(account Account) error {
	var result string
	_, err := time.Parse("2006-01-02", account.Birth)

	if len(account.UserID) < 4 {
		result = "UserID must have at least 4 characters\n"
	}
	if len(account.Password) < 4 {
		result += "Password must have at least 4 characters\n"
	}
	if err != nil {
		result += "Wrong format date"
	}

	if result == "" {
		return nil
	}

	return errors.New(result)
}
