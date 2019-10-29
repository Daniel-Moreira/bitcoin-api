package reports

import (
	. "bitcoin-api-docker/api/src/customtypes"
	"bitcoin-api-docker/api/src/infrastructure/mysql"
	"errors"
	"os"
)

func New(report SystemReport) (map[string]interface{}, error) {
	if report.UserId == "" && report.Date == "" {
		return nil, errors.New("Report must have an userId or a date!")
	}

	command := mysql.SelectCommand{
		TableName:     os.Getenv("TRANSACTIONS_DB"),
		Projection:    []string{"*"},
		Join:          mysql.BITCOIN,
		Conditions:    mysql.USER,
		ConditionData: []string{report.UserId},
	}
	if report.Date != "" {
		command.Conditions = mysql.DATE
		command.ConditionData = []string{report.Date}
	}

	result, err := mysql.Select(command)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"Report": result}, nil
}
