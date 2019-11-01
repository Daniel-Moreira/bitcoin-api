package auth

import (
	. "bitcoin-api-docker/api/src/customtypes"
	"testing"
)

func TestValidate(t *testing.T) {
	account := Account{
		UserID:   "foo4",
		Password: "bar4",
		Birth:    "1992-01-01",
		Name:     "foo",
	}

	err := Validate(account)

	if err != nil {
		t.Error(err.Error())
	}

	account = Account{
		UserID:   "foo",
		Password: "bar4",
		Birth:    "1992-01-01",
		Name:     "foo",
	}

	err = Validate(account)

	if err == nil {
		t.Error("Should fail to validate")
	}

	account = Account{
		UserID:   "foo4",
		Password: "bar",
		Birth:    "1992-01-01",
		Name:     "foo",
	}

	err = Validate(account)

	if err == nil {
		t.Error("Should fail to validate")
	}

	account = Account{
		UserID:   "foo4",
		Password: "bar4",
		Birth:    "gs",
		Name:     "foo",
	}

	err = Validate(account)

	if err == nil {
		t.Error("Should fail to validate")
	}
}
