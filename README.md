# cf-for-k8s-metric-examples

### Examples for Metrics in cf-for-k8s

Setting up metrics in cf-for-k8s involves the following steps:
1. From your app, expose metrics in the Prometheus exposition format
1. In your app's `manifest.yml`, define the appropriate Prometheus annotations

This repo provides examples of applications and the appropriate annotations.

### Deploying an app with metrics

##### Requirements
* cf CLI (7.0.0+)
* a recent [cf-for-k8s deployment](https://github.com/cloudfoundry/cf-for-k8s)
* helm3

##### Deploying the app

Change into your app's root directory and `cf push`

Example for `golang` using cf CLI:
1. `cd go-app-with-metrics`
1. `cf push`

##### Verifying it emits metrics

By defining the Prometheus annotations in the `manifest.yml`, Prometheus will
automatically pick up your app's metrics endpoint.

If you have `kubectl` access on your cluster, you can verify that your app is
emitting metrics by port-forwarding:

```
export POD_NAME="$(k get pods -n cf-workloads | grep go-app-with-metrics | awk '{print $1}')"
export PROM_PORT="YOUR_PORT_HERE"

kubectl port-forward -n cf-workloads $POD_NAME $PROM_PORT

curl localhost:$PROM_PORT/metrics
```

### Deploying Prometheus Server

In order for prometheus to sucessfully scrape pods in a cf-for-k8s cluster,
it currently needs the following:

1. Be deployed in namespace that has the label "istio-injection=enabled".
   This injects istio sidecars onto prometheus's pods. The recommended namespace is `cf-system`
1. Have a [network policy](https://github.com/cloudfoundry/cf-for-k8s-metric-examples/blob/master/prometheus-network-policy.yaml)
   in place that allows prometheus to scrape that namespace.

The network policy requires two new labels:
* a label on the prometheus server's pod: `what-am-i=prometheus`
* a label on the cf-system namespace: `cf-for-k8s.cloudfoundry.org/cf-system-ns: ""`

Using helm3:

* `kubectl apply -f prometheus-network-policy.yaml` (this adds the network
  policy referenced above)
* `kubectl edit namespace cf-system` (and add the label above)
* `helm repo add stable https://kubernetes-charts.storage.googleapis.com`
* `helm install cf-for-k8s-prometheus stable/prometheus -n cf-system --set server.podLabels.what\-am\-i=prometheus`
    * This installs Prometheus in a compatible namespace
    * This adds the label that matches the network policy
* Follow the output to access the Prometheus server

The output should look something like:
```
Get the Prometheus server URL by running these commands in the same shell:
  export POD_NAME=$(kubectl get pods --namespace cf-system -l "app=prometheus,component=server" -o jsonpath="{.items[0].metadata.name}")
  kubectl --namespace cf-system port-forward $POD_NAME 9090
```
* After setting up the port forwarding, access the Prometheus web UI by going to localhost:9090

##### Default Metrics Availability

Metrics should be included for all Prometheus nodes, the API node, and any
pods annotated with Prometheus scrape configurations:

* In a Cloud Foundry manifest:
  ```
  ---
  applications:
  - name: go-app-with-metrics
    metadata:
      annotations:
        prometheus.io/scrape: "true"
        prometheus.io/port: "2112"
        prometheus.io/path: "/metrics"
  ```
* In a Kubernetes pod manifest:
  ```
  spec:
    template:
      metadata:
        annotations:
          prometheus.io/scrape: "true"
          prometheus.io/port: "2112"
          prometheus.io/path: "/metrics"
  ```
