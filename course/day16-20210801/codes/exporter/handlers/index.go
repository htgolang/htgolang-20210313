package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace:   "dev",
		Subsystem:   "web",
		Name:        "http_request_total",
		Help:        "web server http request total",
		ConstLabels: map[string]string{"env": "dev"},
	})

	delayHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "web_process_delay",
		Help:    "Web Process Delay",
		Buckets: prometheus.LinearBuckets(2, 2, 5),
	}, []string{"path"})

	delaySummary = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "web_process_delay_summary",
		Help:       "Web Process Delay Summary",
		Objectives: map[float64]float64{0.6: 0.05, 0.8: 0.02, 0.9: 0.01, 0.95: 0.005, 1: 0.001},
	}, []string{"path"})
)

func InderHandler(w http.ResponseWriter, r *http.Request) {
	delay := float64(rand.Intn(20))
	fmt.Println(delay)
	delayHistogram.WithLabelValues(r.URL.Path).Observe(delay)
	delaySummary.WithLabelValues(r.URL.Path).Observe(delay)

	fmt.Println("InderHandler")
	requestTotal.Inc()
	fmt.Fprintf(w, "%d", time.Now().Unix())
}

func init() {
	prometheus.MustRegister(requestTotal)
	prometheus.MustRegister(delayHistogram)
	prometheus.MustRegister(delaySummary)
}
