package mysql

type InsertCommand struct {
	TableName string
	Data      []interface{}
}

type SelectCommand struct {
	TableName     string
	Projection    []string
	Join          Join
	Conditions    Conditions
	ConditionData []string
}

type Join string
type Conditions string

const (
	BITCOIN Join = "LEFT JOIN users_dev ON users_dev.UserId = transactions_dev.UserId"
	NONE    Join = ""
)

const (
	USER Conditions = "transactions_dev.UserId = ?"
	DATE Conditions = "transactions_dev.Date > ?"
)
