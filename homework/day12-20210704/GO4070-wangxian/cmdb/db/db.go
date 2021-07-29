package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DbConn *sql.DB

func InitDB(dbtype string, dsn string) error {
	var err error
	DbConn, err = sql.Open(dbtype, dsn)
	if err != nil {
		return err
	}

	err = DbConn.Ping()
	if err != nil {
		return err
	}

	return nil
}
