package mysql

type InsertCommand struct {
  TableName string,
  Data []interface{},
}

type SelectCommand struct {
  TableName string,
  Projection []string,
  Conditions Conditions,
  ConditionData []string,
}

type Conditions string

const (
  User Conditions = "userId = ?"
  Day Conditions = "date = ?"
)


