package mysql

import (
	"fmt"
	"strings"

	"github.com/fatih/structs"
	_ "github.com/go-sql-driver/mysql"
)

func (Mysql) Insert(command InsertCommand) error {
	db, err := client()

	defer db.Close()

	if err != nil {
		return err
	}

	data := structs.Map(command.Data[0])
	var keyOrder []string
	columnsString := ""
	for key := range data {
		columnsString += fmt.Sprintf("%s, ", key)
		keyOrder = append(keyOrder, key)
	}
	columnsString = columnsString[0 : len(columnsString)-2]

	var vals []interface{}
	valuesString := ""
	rowString := fmt.Sprintf("(%s),", (strings.Repeat("?, ", len(data)-1) + "?"))
	for i := 0; i < len(command.Data); i++ {
		valuesString += rowString

		rowMap := structs.Map(command.Data[i])
		for i := 0; i < len(keyOrder); i++ {
			vals = append(vals, rowMap[keyOrder[i]])
		}
	}
	valuesString = valuesString[0 : len(valuesString)-1]

	queryString := fmt.Sprintf(
		"INSERT INTO %s(%s) VALUES %s;",
		command.TableName,
		columnsString,
		valuesString,
	)
	stmt, err := db.Prepare(queryString)

	fmt.Println(queryString)
	fmt.Println(vals)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(vals...)

	if err != nil {
		return err
	}

	return nil
}
