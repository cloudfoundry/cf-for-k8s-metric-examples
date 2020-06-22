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

#### Deploy Prometheus Server

using helm3:

* `helm repo add stable https://kubernetes-charts.storage.googleapis.com`
* `helm install <deployment-name> stable/prometheus`
    * This installs Prometheus in the default namespace
* Follow the output to access the Prometheus server

The output should look something like:
```
Get the Prometheus server URL by running these commands in the same shell:
  export POD_NAME=$(kubectl get pods --namespace default -l "app=prometheus,component=server" -o jsonpath="{.items[0].metadata.name}")
  kubectl --namespace default port-forward $POD_NAME 9090
```
* After setting up the port forwarding, access the Prometheus web UI by going to localhost:9090

By default, metrics should be included for all prometheus nodes, the API node,
and any pods annotated with prometheus scrape configurations:
* In a Cloud Foundry manifest:
  ```
  ---
  applications:
  - name: go-app-with-metrics
    buildpacks:
    - https://github.com/cloudfoundry/go-buildpack.git
    metadata:
      annotations:
        prometheus.io/scrape: true
        prometheus.io/port: 2112
        prometheus.io/path: /metrics
  ```
* In a kubernetes pod manifest:
  ```
  spec:
    template:
      metadata:
        annotations:
          prometheus.io/scrape: true
          prometheus.io/port: 2112
          prometheus.io/path: /metrics
  ```

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
