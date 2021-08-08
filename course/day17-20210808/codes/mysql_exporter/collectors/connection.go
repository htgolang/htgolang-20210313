package collectors

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type connectionCollector struct {
	*mysqlCollector
	maxConnectionDesc    *prometheus.Desc
	threadsConnectedDesc *prometheus.Desc
}

func NewConnectionCollector(db *sql.DB) *connectionCollector {
	return &connectionCollector{
		newMySQLCollector(db),
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
	maxConnections := c.variables("max_connections")
	ch <- prometheus.MustNewConstMetric(
		c.maxConnectionDesc,
		prometheus.GaugeValue,
		maxConnections,
	)

	threadsConnected := c.status("threads_connected")
	ch <- prometheus.MustNewConstMetric(
		c.threadsConnectedDesc,
		prometheus.GaugeValue,
		threadsConnected,
	)

	logrus.WithFields(logrus.Fields{
		"threadsConnected": threadsConnected,
		"maxConnections":   maxConnections,
	}).Debug("connection collector")
}
