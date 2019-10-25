package mysql

import (
	"fmt"
	"reflect"
	"strings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type SQLCommand struct {
  TableName string,
  Data []interface{},
}

func Insert(command SQLCommand) {
  // db, err := sql.Open("mysql", "user:pass@tcp(dns)/dbname")
	columns := reflect.ValueOf(command.Data[0]).MapKeys()
  columnsString := strings.Join(columns, ", ")
  
  var vals []interface{}{}
  rowString := fmt.Sprintf("(%s),", (strings.Repeat("?, ", len(columns)-1) + "?"))
  for _, row := range command.Data {
    valuesString += rowString
    for _, v := range row {
      vals = append(vals, v)
    }
  }
  valuesString = valuesString[0:len(valuesString)-1]

  strings.Join(command.Data
	queryString := fmt.Sprintf(
    "INSERT INTO %s(%s) VALUES %s",
    command.tableName,
    columnsString,
    valuesString
  )
	stmt, err := db.Prepare(queryString)

	res, err := stmt.Exec(vals...)
}
