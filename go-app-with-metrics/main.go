package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var customMetric = promauto.NewCounter(prometheus.CounterOpts{
	Name: "app_test_metric",
	Help: "The number of seconds since launch",
})

func incrementMetric() {
	for {
		customMetric.Inc()
		time.Sleep(time.Second)
	}
}

func main() {
	go incrementMetric()

	// Hosting default prometheus handler on :2112/metrics
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
