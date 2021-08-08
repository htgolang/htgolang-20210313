package collectors

import (
	"database/sql"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type trafficCollector struct {
	*mysqlCollector
	desc *prometheus.Desc
}

func NewTrafficCollector(db *sql.DB) *trafficCollector {
	return &trafficCollector{
		newMySQLCollector(db),
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
	directions := []string{"sent", "received"}
	fields := make(logrus.Fields)
	for _, direction := range directions {
		value := c.status(fmt.Sprintf("Bytes_%s", direction))
		ch <- prometheus.MustNewConstMetric(
			c.desc,
			prometheus.CounterValue,
			value,
			direction,
		)
		fields[direction] = value
	}

	logrus.WithFields(fields).Debug("traffic collector")
}
