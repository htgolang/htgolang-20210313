package filters

import (
	"strconv"
	"time"

	"github.com/beego/beego/v2/server/web/context"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	totalRequest = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "cmdb_request_total",
		Help: "cmdb request total",
	})

	urlRequest = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cmdb_request_url_total",
		Help: "cmdb request url total",
	}, []string{"url"})

	statusCode = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cmdb_status_code",
		Help: "cmdb status code",
	}, []string{"status_code"})

	elapsedTime = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "cmdb_url_elapsed_time",
		Help:    "cmdb url elapsed time",
		Buckets: prometheus.LinearBuckets(1, 5, 1),
	}, []string{"url"})
)

func init() {
	prometheus.MustRegister(totalRequest)
	prometheus.MustRegister(urlRequest)
	prometheus.MustRegister(statusCode)
	prometheus.MustRegister(elapsedTime)
}

func BeforeExec(ctx *context.Context) {
	totalRequest.Inc()
	urlRequest.WithLabelValues(ctx.Input.URL()).Inc()
	ctx.Input.SetData("stime", time.Now().Unix())
}

func AfterExec(ctx *context.Context) {
	statusCode.WithLabelValues(strconv.Itoa(ctx.ResponseWriter.Status)).Inc()
	stime := ctx.Input.GetData("stime")
	if stime != nil {
		if st, ok := stime.(int64); ok {
			elapsed := time.Now().Unix() - st
			elapsedTime.WithLabelValues(ctx.Input.URL()).Observe(float64(elapsed))
		}
	}
}
