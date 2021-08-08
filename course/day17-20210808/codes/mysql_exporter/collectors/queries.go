package collectors

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type slowQueriesCollector struct {
	*mysqlCollector
	desc *prometheus.Desc
}

func NewSlowQueriesCollector(db *sql.DB) *slowQueriesCollector {
	return &slowQueriesCollector{
		newMySQLCollector(db),
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
	value := c.status("slow_queries")
	ch <- prometheus.MustNewConstMetric(
		c.desc,
		prometheus.CounterValue,
		value,
	)

	logrus.WithFields(logrus.Fields{
		"metric": value,
	}).Debug("slow queries collector")
}

type qpsCollector struct {
	*mysqlCollector
	desc *prometheus.Desc
}

func NewQpsCollector(db *sql.DB) *qpsCollector {
	return &qpsCollector{
		newMySQLCollector(db),
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
	value := c.status("Queries")

	ch <- prometheus.MustNewConstMetric(
		c.desc,
		prometheus.CounterValue,
		value,
	)

	logrus.WithFields(logrus.Fields{
		"metric": value,
	}).Debug("qps collector")
}
