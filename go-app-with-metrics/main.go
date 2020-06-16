package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Hosting default prometheus handler on :2112/metrics
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
