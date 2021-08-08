package collectors

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
)

type commandCollector struct {
	db   *sql.DB
	desc *prometheus.Desc
}

func NewCommandCollector(db *sql.DB) *commandCollector {
	return &commandCollector{
		db,
		prometheus.NewDesc(
			"mysql_global_status_command",
			"mysql global status command",
			[]string{"command"},
			nil,
		),
	}
}

func (c *commandCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.desc
}

func (c *commandCollector) Collect(ch chan<- prometheus.Metric) {
	variableName, variableValue := "", 0

	sql := "show global status where variable_name = 'Com_insert'"
	value := 0

	if err := c.db.QueryRow(sql).Scan(&variableName, &variableValue); err == nil {
		value = variableValue
	}

	ch <- prometheus.MustNewConstMetric(
		c.desc,
		prometheus.CounterValue,
		float64(value),
		"insert",
	)

	sql = "show global status where variable_name = 'Com_update'"
	value = 0

	if err := c.db.QueryRow(sql).Scan(&variableName, &variableValue); err == nil {
		value = variableValue
	}

	ch <- prometheus.MustNewConstMetric(
		c.desc,
		prometheus.CounterValue,
		float64(value),
		"update",
	)

	sql = "show global status where variable_name = 'Com_select'"
	value = 0

	if err := c.db.QueryRow(sql).Scan(&variableName, &variableValue); err == nil {
		value = variableValue
	}

	ch <- prometheus.MustNewConstMetric(
		c.desc,
		prometheus.CounterValue,
		float64(value),
		"select",
	)

	sql = "show global status where variable_name = 'Com_delete'"
	value = 0

	if err := c.db.QueryRow(sql).Scan(&variableName, &variableValue); err == nil {
		value = variableValue
	}

	ch <- prometheus.MustNewConstMetric(
		c.desc,
		prometheus.CounterValue,
		float64(value),
		"delete",
	)
}
