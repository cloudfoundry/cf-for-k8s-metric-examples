package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main() {
	completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_last_completion_timestamp_seconds",
		Help: "The timestamp of the last successful completion of a DB backup.",
	})
	err := push.New("pushgateway-test-prometheus-pushgateway.cf-system.svc.cluster.local:9091", "test").Collector(completionTime).Push()
	fmt.Printf("error: %s", err)
}
