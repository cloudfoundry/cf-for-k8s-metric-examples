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
   This injects istio sidecars onto prometheus's pods. The recommended namespace is `cf-system`.
1. Have a network policy in place that allows prometheus to scrape your app's
   pod.
1. Have the necessary certs available in prometheus's istio sidecar.

Because of the complexity of these requirements, we recommend using
[cf-k8s-prometheus](https://github.com/cloudfoundry/cf-k8s-prometheus#how-to-deploy-in-cf-for-k8s)

* Follow the output to access the Prometheus server

The output should look something like:
```
Get the Prometheus server URL by running these commands in the same shell:
  export POD_NAME=$(kubectl get pods --namespace cf-system -l "metrics=prometheus,component=server" -o jsonpath="{.items[0].metadata.name}")
  kubectl --namespace cf-system port-forward $POD_NAME 9090
```
* After setting up the port forwarding, access the Prometheus web UI by going to localhost:9090

##### Default Metrics Availability

Metric sshould be included for all Prometheus nodes, the API node, and any
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
