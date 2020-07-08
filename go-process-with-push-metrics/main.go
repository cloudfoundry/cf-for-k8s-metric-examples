package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main() {
	pushedMetric := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "pushed_metric",
		Help: "Testing that a metric can be pushed to the gateway.",
	})

	pushedMetric.Add(1)

	err := push.New("localhost:9091", "pushed-metrics").Collector(pushedMetric).Push()
	if err != nil {
		fmt.Printf("error: %s", err)
		panic(err)
	}
}
