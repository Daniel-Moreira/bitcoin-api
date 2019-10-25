package mysql

type InsertCommand struct {
  TableName string,
  Data []interface{},
}

type SelectCommand struct {
  TableName string,
  Projection []string,
  Join Join
  Conditions Conditions,
  ConditionData []string,
}

type Join string
type Conditions string

const (
  BitCoin Join = "RIGHT JOIN coins ON user.userId = coins.userId"
  None Join = ""
)

const (
  User Conditions = "userId = ?"
  Day Conditions = "date = ?"
)
