package mysql

import (
	. "bitcoin-api-docker/api/src/customtypes"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func UpdateCoinAmount(tableName string, transaction Transaction) error {
	db, err := client()

	defer db.Close()

	if err != nil {
		return err
	}

	queryString := fmt.Sprintf(
		"UPDATE %s SET BitCoinAmount = BitCoinAmount + ? WHERE UserId = ? AND BitCoinAmount + ? > 0;",
		tableName,
	)

	stmt, err := db.Prepare(queryString)

	fmt.Println(queryString)
	if err != nil {
		return err
	}

	amount := transaction.Amount
	if transaction.Type == "sell" {
		amount = "-" + amount
	}
	result, err := stmt.Exec(amount, transaction.UserId, amount)

	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Not enought bitcoins")
	}

	return nil
}
