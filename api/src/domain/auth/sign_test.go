package auth

import (
	. "bitcoin-api-docker/api/src/customtypes"
	"bitcoin-api-docker/api/src/infrastructure/mysql"
	"testing"
)

type MockMysql struct{}

func (MockMysql) Insert(command mysql.InsertCommand) error {
	return nil
}

func TestSign(t *testing.T) {
	account := Account{
		UserID:   "fooo",
		Password: "barr",
		Birth:    "1992-01-01",
		Name:     "foo",
	}

	mockSql := &MockMysql{}
	_, err := Sign(account, mockSql)

	if err != nil {
		t.Error(err.Error())
	}
}
