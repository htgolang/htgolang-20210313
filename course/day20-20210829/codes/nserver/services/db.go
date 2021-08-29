package services

import (
	"database/sql"
)

var Db *sql.DB

func InitDb(driverName, dsn string) error {
	var err error
	Db, err = sql.Open(driverName, dsn)
	if err != nil {
		return err
	}
	err = Db.Ping()
	if err != nil {
		return err
	}
	return nil
}
