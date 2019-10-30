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
	BITCOIN Join = "INNER JOIN users_dev USING(UserId)"
	NONE    Join = ""
)

const (
	USER Conditions = "UserId = ?"
	DATE Conditions = "Date(transactions_dev.Date) = ?"
)
