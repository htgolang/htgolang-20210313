package collectors

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
)

type trafficCollector struct {
	db   *sql.DB
	desc *prometheus.Desc
}

func NewTrafficCollector(db *sql.DB) *trafficCollector {
	return &trafficCollector{
		db,
		prometheus.NewDesc(
			"mysql_global_status_traffic",
			"mysql global status traffic",
			[]string{"direction"},
			nil,
		),
	}
}

func (c *trafficCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.desc
}

func (c *trafficCollector) Collect(ch chan<- prometheus.Metric) {
	sql := "show global status where variable_name='Bytes_sent'"
	value := 0

	variableName, variableValue := "", 0
	if err := c.db.QueryRow(sql).Scan(&variableName, &variableValue); err == nil {
		value = variableValue
	}

	ch <- prometheus.MustNewConstMetric(
		c.desc,
		prometheus.CounterValue,
		float64(value),
		"sent",
	)

	sql = "show global status where variable_name='Bytes_received'"

	if err := c.db.QueryRow(sql).Scan(&variableName, &variableValue); err == nil {
		value = variableValue
	}

	ch <- prometheus.MustNewConstMetric(
		c.desc,
		prometheus.CounterValue,
		float64(value),
		"recevied",
	)

}
