package collectors

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
)

type slowQueriesCollector struct {
	db   *sql.DB
	desc *prometheus.Desc
}

func NewSlowQueriesCollector(db *sql.DB) *slowQueriesCollector {
	return &slowQueriesCollector{
		db,
		prometheus.NewDesc(
			"mysql_global_status_slow_queries",
			"mysql global status slow queries",
			nil,
			nil,
		),
	}
}

func (c *slowQueriesCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.desc
}

func (c *slowQueriesCollector) Collect(ch chan<- prometheus.Metric) {
	value := 0

	variableName, variableValue := "", 0

	sql := "show global status where variable_name='slow_queries'"
	if err := c.db.QueryRow(sql).Scan(&variableName, &variableValue); err == nil {
		value = variableValue
	}

	ch <- prometheus.MustNewConstMetric(
		c.desc,
		prometheus.CounterValue,
		float64(value),
	)
}

type qpsCollector struct {
	db   *sql.DB
	desc *prometheus.Desc
}

func NewQpsCollector(db *sql.DB) *qpsCollector {
	return &qpsCollector{
		db,
		prometheus.NewDesc(
			"mysql_global_status_queries",
			"mysql global status queries",
			nil,
			nil,
		),
	}
}

func (c *qpsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.desc
}

func (c *qpsCollector) Collect(ch chan<- prometheus.Metric) {
	value := 0

	variableName, variableValue := "", 0

	sql := "show global status where variable_name='Queries'"
	if err := c.db.QueryRow(sql).Scan(&variableName, &variableValue); err == nil {
		value = variableValue
	}

	ch <- prometheus.MustNewConstMetric(
		c.desc,
		prometheus.CounterValue,
		float64(value),
	)
}
