package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"exporter/collectors"
	"exporter/handlers"
	"exporter/services"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	rand.Seed(time.Now().Unix())

	services.StartSyncTask()
	services.StartCpuStat()

	// 1. 定义指标
	prometheus.MustRegister(prometheus.NewCounterFunc(prometheus.CounterOpts{
		Name: "web_server_current_time",
		Help: "Web Server Current Time",
	}, func() float64 {
		fmt.Println("current Time")
		return (float64)(time.Now().Unix())
	}))

	prometheus.MustRegister(collectors.NewMemPercentCollector())

	addr := ":9999"
	http.Handle("/", handlers.AuthHandler(http.HandlerFunc(handlers.InderHandler)))
	http.Handle("/metrics/", handlers.AuthHandler(promhttp.Handler()))
	http.ListenAndServe(addr, nil)
}
