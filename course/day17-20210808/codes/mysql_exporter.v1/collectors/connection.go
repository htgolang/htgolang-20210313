package collectors

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
)

type connectionCollector struct {
	db                   *sql.DB
	maxConnectionDesc    *prometheus.Desc
	threadsConnectedDesc *prometheus.Desc
}

func NewConnectionCollector(db *sql.DB) *connectionCollector {
	return &connectionCollector{
		db,
		prometheus.NewDesc(
			"mysql_global_variables_max_connections",
			"mysql global variables max connections",
			nil,
			nil,
		),
		prometheus.NewDesc(
			"mysql_global_status_threads_connected",
			"mysql global status threads connected",
			nil,
			nil,
		),
	}
}

func (c *connectionCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.maxConnectionDesc
	ch <- c.threadsConnectedDesc
}

func (c *connectionCollector) Collect(ch chan<- prometheus.Metric) {
	value := 0
	variableName, variableValue := "", 0
	sql := "show global variables where variable_name='max_connections'"

	if err := c.db.QueryRow(sql).Scan(&variableName, &variableValue); err == nil {
		value = variableValue
	}

	ch <- prometheus.MustNewConstMetric(
		c.maxConnectionDesc,
		prometheus.GaugeValue,
		float64(value),
	)

	value = 0
	sql = "show global status where variable_name='threads_connected'"

	if err := c.db.QueryRow(sql).Scan(&variableName, &variableValue); err == nil {
		value = variableValue
	}

	ch <- prometheus.MustNewConstMetric(
		c.threadsConnectedDesc,
		prometheus.GaugeValue,
		float64(value),
	)
}
