package auth

import (
	"os"

	. "bitcoin-api-docker/api/src/customtypes"
	"bitcoin-api-docker/api/src/infrastructure/mysql"
)

func Sign(account Account) (map[string]string, error) {
	err := Validate(account)

	if err != nil {
		return nil, err
	}

	command := mysql.InsertCommand{
		TableName: os.Getenv("USERS_DB"),
	}
	command.Data = append(command.Data, account)

	err = mysql.Insert(command)

	if err != nil {
		return nil, err
	}

	return map[string]string{"Message": "Account Created!"}, nil
}
