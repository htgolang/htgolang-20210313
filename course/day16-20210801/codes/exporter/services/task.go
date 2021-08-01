package services

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	syncTaskTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "sync_task_total",
		Help: "Sync data task total",
	}, []string{"type"})
)

func StartSyncTask() {
	// 同步任务1
	go func() {
		for range time.Tick(time.Second * 10) {
			fmt.Println("task1")
			syncTaskTotal.WithLabelValues("task1").Inc()
		}
	}()
	// 同步任务2
	go func() {
		for range time.Tick(time.Second * 3) {
			fmt.Println("task2")
			syncTaskTotal.WithLabelValues("task2").Inc()
		}
	}()
}

func init() {
	prometheus.MustRegister(syncTaskTotal)
}
