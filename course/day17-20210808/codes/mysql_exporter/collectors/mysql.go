package collectors

import "database/sql"

type mysqlCollector struct {
	db *sql.DB
}

func newMySQLCollector(db *sql.DB) *mysqlCollector {
	return &mysqlCollector{db}
}

func (c *mysqlCollector) status(key string) float64 {
	sql := "show global status where variable_name = ?"
	name, value := "", 0.0

	if err := c.db.QueryRow(sql, key).Scan(&name, &value); err == nil {
		return value
	}

	return 0.0
}

func (c *mysqlCollector) variables(key string) float64 {
	sql := "show global variables where variable_name = ?"
	name, value := "", 0.0

	if err := c.db.QueryRow(sql, key).Scan(&name, &value); err == nil {
		return value
	}

	return 0.0
}
