package mysql

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func Select(command SelectCommand) ([]map[string]string, error) {
	db, err := client()

	defer db.Close()

	if err != nil {
		return nil, err
	}

	projectionString := strings.Join(command.Projection, ", ")

	queryString := fmt.Sprintf(
		"SELECT %s FROM %s %s WHERE %s;",
		projectionString,
		command.TableName,
		command.Join,
		command.Conditions,
	)

	data := command.ConditionData
	dataInterface := make([]interface{}, len(data))
	for i, v := range data {
		dataInterface[i] = v
	}

	rows, err := db.Query(queryString, dataInterface...)

	if err != nil {
		return nil, err
	}

	columnNames, err := rows.Columns()
	columnSize := len(columnNames)

	rawRow := make([]interface{}, columnSize)
	for index := range rawRow {
		rawRow[index] = &rawRow[index]
	}

	result := make([]map[string]string, 0)
	obj := map[string]string{}
	for rows.Next() {
		err = rows.Scan(rawRow...)

		if err != nil {
			return nil, err
		}

		for i := 0; i < columnSize; i++ {
			obj[columnNames[i]] = fmt.Sprintf("%v", string(rawRow[i].([]byte)))
		}

		result = append(result, obj)
	}

	return result, nil
}
