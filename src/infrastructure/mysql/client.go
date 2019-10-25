package mysql

import (
  "os"
  "database/sql"

  _ "github.com/go-sql-driver/mysql"
)

func client () (*sql.DB, error) {
  db, err := sql.Open(
    "mysql",
    fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
      os.Getenv("MYSQL_USER"),
      os.Getenv("MYSQL_PASSWORD"),
      os.Getenv("MYSQL_HOST"),
      os.Getenv("MYSQL_PORT"),
      os.Getenv("MYSQL_DATABASE")
    )
  )

  return db, err
}
