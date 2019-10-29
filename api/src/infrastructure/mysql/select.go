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

	fmt.Println(queryString)

	data := command.ConditionData
	dataInterface := make([]interface{}, len(data))
	for i, v := range data {
		dataInterface[i] = v
	}

	fmt.Println(data)
	rows, err := db.Query(queryString, dataInterface...)

	if err != nil {
		return nil, err
	}

	columnNames, err := rows.Columns()
	columnSize := len(columnNames)

	rawRow := make([]interface{}, columnSize)

	result := make([]map[string]string, 0)
	for rows.Next() {
		obj := map[string]string{}
		for index := range rawRow {
			rawRow[index] = &rawRow[index]
		}
		err = rows.Scan(rawRow...)

		if err != nil {
			return nil, err
		}

		for i := 0; i < columnSize; i++ {
			switch rawRow[i].(type) {
			case int64:
				obj[columnNames[i]] = fmt.Sprintf("%d", rawRow[i])
			case float64:
				obj[columnNames[i]] = fmt.Sprintf("%f", rawRow[i])
			case float32:
				obj[columnNames[i]] = fmt.Sprintf("%f", rawRow[i])
			default:
				obj[columnNames[i]] = fmt.Sprintf("%v", string(rawRow[i].([]byte)))
			}
		}

		result = append(result, obj)
	}

	fmt.Println(result)

	return result, nil
}
