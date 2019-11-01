package auth

import (
	"os"

	. "bitcoin-api-docker/api/src/customtypes"
	. "bitcoin-api-docker/api/src/infrastructure/mysql"
)

func Sign(account Account, sql Sql) (map[string]string, error) {
	err := Validate(account)

	if err != nil {
		return nil, err
	}

	command := InsertCommand{
		TableName: os.Getenv("USERS_DB"),
	}
	command.Data = append(command.Data, account)

	err = sql.Insert(command)

	if err != nil {
		return nil, err
	}

	return map[string]string{"Message": "Account Created!"}, nil
}
