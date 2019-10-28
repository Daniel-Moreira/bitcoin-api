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
	BITCOIN Join = "RIGHT JOIN coins ON user.userId = coins.userId"
	NONE    Join = ""
)

const (
	USER Conditions = "userId = ?"
	DAY  Conditions = "date = ?"
)
