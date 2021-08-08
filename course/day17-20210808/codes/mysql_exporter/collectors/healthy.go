package collectors

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type healthyCollector struct {
	db   *sql.DB
	desc *prometheus.Desc
}

func NewHealthyCollector(db *sql.DB) *healthyCollector {
	return &healthyCollector{
		db: db,
		desc: prometheus.NewDesc(
			"mysql_healthy",
			"MySQL Healthy Status",
			nil,
			nil,
		),
	}
}

func (c *healthyCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.desc
}

func (c *healthyCollector) Collect(ch chan<- prometheus.Metric) {
	isHealthy := 1
	if err := c.db.Ping(); err != nil {
		isHealthy = 0
	}

	ch <- prometheus.MustNewConstMetric(
		c.desc,
		prometheus.GaugeValue,
		float64(isHealthy),
	)

	logrus.WithFields(logrus.Fields{
		"isHealthy": isHealthy,
	}).Debug("healthy collector")
}
