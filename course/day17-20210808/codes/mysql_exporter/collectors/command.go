package collectors

import (
	"database/sql"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/sirupsen/logrus"
)

type commandCollector struct {
	*mysqlCollector
	desc *prometheus.Desc
}

func NewCommandCollector(db *sql.DB) *commandCollector {
	return &commandCollector{
		newMySQLCollector(db),
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
	commands := []string{"insert", "update", "delete", "select"}

	for _, cmd := range commands {
		value := c.status(fmt.Sprintf("Com_%s", cmd))
		ch <- prometheus.MustNewConstMetric(
			c.desc,
			prometheus.CounterValue,
			value,
			cmd,
		)
		logrus.WithFields(logrus.Fields{
			"command": cmd,
			"metric":  value,
		}).Debug("command collector")
	}
}
