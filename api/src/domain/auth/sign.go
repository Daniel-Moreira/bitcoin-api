package auth

import (
	"os"

	"bitcoin-api/src/infrastructure/mysql"
	. "bitcoin-api/src/customtypes"
)

func Sign(account Account) (map[string]string, error) {
	err := Validate(account)

	if err != nil {
		return nil, err
	}

  command := mysql.InsertCommand{
    TableName: os.Getenv("REGISTER_USERS"),
    Data: account
  }

	err = mysql.insert(command)

	if err != nil {
		return nil, err
	}

	return map[string]string{"Message": "Account Created!"}, nil
}
