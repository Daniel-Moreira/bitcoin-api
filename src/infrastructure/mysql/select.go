package mysql

import (
	"fmt"
	"reflect"
  "strings"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Select(command SelectCommand) (error) {
  db, err := client()

  defer db.Close()

  if err != nil {
    return err
  }

  projectionString := strings.Join(command.Projection, ", ")

	queryString := fmt.Sprintf(
    "SELECT %s FROM %s WHERE %s;",
    projectionString,
    command.TableName,
    command.Conditions
  )
  rows, err := db.Query(queryString, ConditionData...)

  if err != nil {
    return err
  }

  columnNames, err := rows.Columns()
  columnSize := len(columnNames)

  rawRow := make([]interface{}, columnSize)
  var result []map[string][string]
  var obj map[string][string]
  for rows.Next() {
    err = rows.Scan(rawRow...)

    if err != nil {
      return err
    }

    for i := 0; i < columnSize; i++ {
      obj[columnNames[i]] = string(rawRow[i])
    }

    result = append(result, obj)
  }

  return result
}
