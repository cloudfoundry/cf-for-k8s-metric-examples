package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var customMetric = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "app_test_metric",
	Help: "The number of seconds since launch",
}, []string{"query"})

func incrementMetric() {
	for {
		customMetric.WithLabelValues("").Inc()
		time.Sleep(time.Second)
	}
}

func main() {
	go incrementMetric()

	// Hosting default prometheus handler on :2112/metrics
	http.Handle("/metrics", queryParamWrapper(promhttp.Handler()))
	http.ListenAndServe(":2112", nil)
}

func queryParamWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, _ := range r.URL.Query() {
			customMetric.WithLabelValues(k).Inc()
		}
		h.ServeHTTP(w, r)
	})
}
