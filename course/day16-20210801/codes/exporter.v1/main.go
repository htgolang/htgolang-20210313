package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

var (
	requestTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace:   "dev",
		Subsystem:   "web",
		Name:        "http_request_total",
		Help:        "web server http request total",
		ConstLabels: map[string]string{"env": "dev"},
	})

	syncTaskTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "sync_task_total",
		Help: "Sync data task total",
	}, []string{"type"})

	currentTime = prometheus.NewCounterFunc(prometheus.CounterOpts{
		Name: "web_server_current_time",
		Help: "Web Server Current Time",
	}, func() float64 {
		fmt.Println("current Time")
		return (float64)(time.Now().Unix())
	})

	cpuPercent = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "web_server_cpu_percent",
		Help: "Web Server CPU Percent",
	}, []string{"cpu"})

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

func validateUser(txt string) bool {
	pos := strings.Index(txt, ":")
	if pos < 0 {
		return false
	}
	user, password := txt[:pos], txt[pos+1:]
	if pwd, ok := users[user]; ok {
		// pwd // MD5(配置/数据)
		return pwd == fmt.Sprintf("%X", md5.Sum([]byte(password)))
	}
	return false
}

func AuthHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 检查前
		authorization := r.Header.Get("Authorization")
		prefix := "Basic "
		if strings.HasPrefix(authorization, prefix) {
			authorization = authorization[len(prefix):]
		}
		txt, err := base64.StdEncoding.DecodeString(authorization)
		if err != nil || !validateUser(string(txt)) {
			// 认证失败
			w.Header().Add("www-authenticate", "Basic")
			w.WriteHeader(401)
			return
		}

		handler.ServeHTTP(w, r)
		// 检查后
	})
}

func main() {
	rand.Seed(time.Now().Unix())
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

	go func() {
		for range time.Tick(time.Second * 10) {
			for i := 0; i < 4; i++ {
				fmt.Println("cpu percent", i)
				cpuPercent.WithLabelValues(strconv.Itoa(i)).Set(rand.Float64() * 100)
			}
		}
	}()

	// 1. 定义指标
	prometheus.MustRegister(requestTotal)
	prometheus.MustRegister(syncTaskTotal)
	prometheus.MustRegister(currentTime)
	prometheus.MustRegister(cpuPercent)
	prometheus.MustRegister(delayHistogram)
	prometheus.MustRegister(delaySummary)
	prometheus.MustRegister(NewMemPercentCollector())
	// 2. 注册指标
	// 3. 注册处理器
	// 4. 启动web服务
	// 5. 更新指标信息

	// fmt.Println(prometheus.LinearBuckets(2, 2, 5))

	addr := ":9999"
	http.Handle("/", AuthHandler(http.HandlerFunc(InderHandler)))
	http.Handle("/metrics/", AuthHandler(promhttp.Handler()))
	http.ListenAndServe(addr, nil)
}
