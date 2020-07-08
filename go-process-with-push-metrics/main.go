package main

import (
	"fmt"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main() {
	pushedMetric := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "pushed_metric",
		Help: "Testing that a metric can be pushed to the gateway.",
	})

	pushedMetric.Add(1)

	pushGatewayAddr := os.Getenv("PUSHGATEWAY_ADDR")
	if pushGatewayAddr == "" {
		pushGatewayAddr = "localhost:9091"
	}

	err := push.New(pushGatewayAddr, "pushed-metrics").Collector(pushedMetric).Push()
	if err != nil {
		fmt.Printf("error: %s", err)
		panic(err)
	}
}
