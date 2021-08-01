package services

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	cpuPercent = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "web_server_cpu_percent",
		Help: "Web Server CPU Percent",
	}, []string{"cpu"})
)

func StartCpuStat() {
	go func() {
		for range time.Tick(time.Second * 10) {
			for i := 0; i < 4; i++ {
				fmt.Println("cpu percent", i)
				cpuPercent.WithLabelValues(strconv.Itoa(i)).Set(rand.Float64() * 100)
			}
		}
	}()

}

func init() {
	prometheus.MustRegister(cpuPercent)
}
