# cf-for-k8s-metric-examples

### Examples for Metrics in cf-for-k8s

Setting up metrics in cf-for-k8s involves two steps:
1. From your app, expose metrics in the Prometheus exposition format
2. In your app's `manifest.yml`, define the appropriate Prometheus annotations

This repo provides examples of applications and the appropriate annotations.

### Deploying an app with metrics

##### Deploying the app

Change into your app's root directory and `cf push`

Example for `golang`:
1. `cd go-app-with-metrics`
2. `cf push`

##### Verifying it emits metrics

By defining the Prometheus annotations in the `manifest.yml`, Prometheus will
automatically pick up your app's metrics endpoint.

If you have `kubectl` access on your cluster, you can verify that your app is
emitting metrics by port-forwarding:

```
export POD_NAME="YOUR_POD_NAME_HERE"
export PROM_PORT="YOUR_PORT_HERE"
kubectl port-forward -n cf-workloads $POD_NAME $PROM_PORT

curl localhost:$PROM_PORT/metrics
```
