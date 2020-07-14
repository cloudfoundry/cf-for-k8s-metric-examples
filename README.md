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
1. `cf7 push`


### Deploying Prometheus Server

Using helm3:


* `helm repo add stable https://kubernetes-charts.storage.googleapis.com`
* `helm install cf-for-k8s-prometheus stable/prometheus -n cf-system --set server.podLabels.what\-am\-i=prometheus`
    * This installs Prometheus in a compatable namespace and label that matches the network policy
* `kubectl apply -f prometheus-network-policy.yaml` (this adds the network

* Follow the output to access the Prometheus server

The output should look something like:
```
Get the Prometheus server URL by running these commands in the same shell:
  export POD_NAME=$(kubectl get pods --namespace cf-system -l "app=prometheus,component=server" -o jsonpath="{.items[0].metadata.name}")
  kubectl --namespace cf-system port-forward $POD_NAME 9090
```
* After setting up the port forwarding, access the Prometheus web UI by going to localhost:9090

##### Verifying it emits metrics

By defining the Prometheus annotations in the `manifest.yml`, Prometheus will
automatically pick up your app's metrics endpoint. Once you have prometheus
deployed and your app annotations in place you will be able to verify a list of
metrics by executing the following query `{kubernetes_namespace="cf-workloads"}`.

##### Default Metrics Availability
By configuring installing prometheus from the helm chart above, and applying
the network policy Prometheus Server will include the following metrics.

* Annotated cf-for-k8s-component metrics (currently includes CAPI, UAA, Log Cache, Metric Proxy)
* Annotated cf pushed apps.
* Container Metrics from [metrics-server](https://github.com/kubernetes-sigs/metrics-server)
* Node Metrics from [node exporter](https://github.com/prometheus/node_exporter)
