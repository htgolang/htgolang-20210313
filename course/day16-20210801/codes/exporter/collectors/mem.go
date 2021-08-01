package collectors

import (
	"fmt"
	"math/rand"

	"github.com/prometheus/client_golang/prometheus"
)

type MemPercentCollector struct {
	desc *prometheus.Desc
}

func NewMemPercentCollector() *MemPercentCollector {
	return &MemPercentCollector{
		prometheus.NewDesc("mem_percent", "Mem Percent", nil, nil),
	}
}

func (c *MemPercentCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.desc
}

func (c *MemPercentCollector) Collect(ch chan<- prometheus.Metric) {
	// 四种类型
	fmt.Println("mem Percent")
	ch <- prometheus.MustNewConstMetric(c.desc, prometheus.GaugeValue, rand.Float64()*100)
}
