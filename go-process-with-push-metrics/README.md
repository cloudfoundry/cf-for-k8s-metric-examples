### go-process-with-push-metrics

This is an example application that pushes a counter metric into a local
pushgateway.


#### How to Run

Get the latest [Prometheus and Pushgateway](https://prometheus.io/download),
using the included `prometheus.yml`, run both the Prometheus server and
Pushgateway.

Once both are up, push metrics to the Pushgateway by running this app using
`go run main.go`. Then you can verify metrics are being served by
the Pushgateway, and ultimately scraped by the Prometheus server.

Example:
```
# Download and Start Prometheus using our config file
wget https://github.com/prometheus/prometheus/releases/download/v2.19.2/prometheus-2.19.2.linux-amd64.tar.gz
tar -xvf prometheus-2.19.2.linux-amd64.tar.gz
./prometheus-2.19.2.linux-amd64/prometheus --config.file=./prometheus.yml &
export PROM_PID=$!

# Download and Start Push Gateway
wget https://github.com/prometheus/pushgateway/releases/download/v1.2.0/pushgateway-1.2.0.linux-amd64.tar.gz
tar -xvf pushgateway-1.2.0.linux-amd64.tar.gz
./pushgateway-1.2.0.linux-amd64/pushgateway &
export PUSH_PID=$!

# Rush metrics to the pushgateway
go run main.go

# Verify that the metric was pushed to Pushgateway
curl localhost:9091/metrics

# Verify that the metric was scraped to the Prometheus server
#   (after waiting for scrape interval)
sleep 15
curl localhost:9090/api/v1/query?query=pushed_metric

### Optional cleanup
kill $PROM_PID
rm prometheus-2.19.2.linux-amd64.tar.gz*
rm -rf prometheus-2.19.2.linux-amd64

kill $PUSH_PID
rm pushgateway-1.2.0.linux-amd64.tar.gz*
rm -rf pushgateway-1.2.0.linux-amd64
```
