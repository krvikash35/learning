
## Workload Resources/Objects
* `kubectl api-resources` return the list of supported object
* types
  * Deployment
  * Pod
  * Node
  * Service
  * Ingress
  * ..etc

## service
* it is used to expose the service
* K8s relies on proxying to forward inbound traffic to backend rather than using multi-ip dns for round-robin resolution. It is proxied via kube-proxy that run on each node. 
* kube-proxy in ip-table mode select random backend &  has low overhead because traffic is handled at linux low level. while in user mode, it can check if pod is healthy else direct traffic to other backend pod.
* type
    * clusterIP: Exposes the service on a cluster-internal IP. Choosing this value makes the service only reachable from within the cluster. This is the default ServiceType.
    * NodePort: extension of clusterIP. Exposes the service on each Node’s IP at a static port (the NodePort). A ClusterIP service, to which the NodePort service will route, is automatically created. You’ll be able to contact the NodePort service, from outside the cluster, by requesting `<NodeIP>:<NodePort>`.
    * LoadBalancer: extension of Nodeport.  Exposes the service externally using a cloud provider’s(gcp, aws etc) load balancer. NodePort and ClusterIP services, to which the external load balancer will route, are automatically created.
* k8s create dns record for service/pod. we can use service name instead of ip to access it.
* default fqdn: `<service-name>.<namespace>.svc.cluster.local`. No need to use fqdn to access service in same namespace. i.e `nslookup simple-node-server.default.svc.cluster.local` will return cluster ip of service.
![](./service.png)

### NodePort
* `port`: is used when one service wants to talk to other service.
* `nodePort`: is used when someone wants to talk to service outside cluster. Control plane will auto assign port from available ports(default: 30000-32767). We can assign manually but it might lead to collision if same port is already allocated. traffic--->nodePort-->port-->targetPort
* `targetPort`: it is actual port that service is listening on inside container.
```
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  type: NodePort
  selector:
    app.kubernetes.io/name: MyApp
  ports:
    - port: 80
      targetPort: 80
      nodePort: 30007
```

### LoadBalancer
*  The cloud provider decides how it is load balanced. The cloud-controller-manager component then configures the external load balancer to forward traffic to that assigned node port.
* If you specify a loadBalancerIP but your cloud provider does not support the feature, the loadbalancerIP field that you set is ignored.
```
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  selector:
    app.kubernetes.io/name: MyApp
  ports:
    - protocol: TCP
      port: 80
      targetPort: 9376
  clusterIP: 10.0.171.239
  type: LoadBalancer
status:
  loadBalancer:
    ingress:
    - ip: 192.0.2.127
```

## Ingress
* LoadBalancer service type can be expensive to manage and does not support url based routing, ssl termination etc.
* ingress can also be used to expose service under same ip address using different name. mainly used for http service.
* You must have an Ingress controller( i.e `ngnix-ingress`) to satisfy an Ingress. Only creating an Ingress resource has no effect.
* it does not expose arbitrary ports or protocol. To expose services than http/https(i.e may be grpc), we require servcie `NodePort` or `LoadBalancer`.

Ingress controller: 
* it can be software load balancer or hardware or cloud load balancer running externally. Each one have its own feature advantage.
* It has to run as pod, only expose 80/443 on externally. always use ssl at external LB.
* examples: Nginx, Ambassador, EnRoute, HAProxy, AWS ALB, AKS Application Gateway

```
kind: Ingress
  apiVersion: extensions/v1beta1
  metadata:
    name: "example-ingress"
    namespace: production
    annotations:
      kubernetes.io/ingress.class: alb
      alb.ingress.kubernetes.io/scheme: internet-facing
      alb.ingress.kubernetes.io/listen-ports: '[{"HTTPS":443}, {"HTTP":80}]'
  spec:
    rules:
      - host: abc.com
        http:
          paths:
            - path: /orders
              backend:
                serviceName: order-service
                servicePort: 8080
            - path: /blog
              backend:
                serviceName: blog-service
                servicePort: 8081
```



![](./ingress.png)

## Horizontal pod autoscalling
* we can enable hpa based on resource utilization(i.e cpu usage).
* we have two option, use kubectl cli or resource manifest file.

`kubectl autoscale deployment php-apache --cpu-percent=50 --min=1 --max=10`

```
apiVersion: autoscaling/v2beta2
kind: HorizontalPodAutoscaler
metadata:
  name: my-app-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: my-app-deployment
  minReplicas: 1
  maxReplicas: 3
  metrics:
    - type: Resource
      resource:
        name: cpu
        target:
          type: Utilization
          averageUtilization: 50
```
## commands

* `kubectl create` is imperative while `kubectl apply` is declarative pattern. apply will not throw error if resource is already created.
```
kubectl cluster-info
kubectl get pods
kubectl get nodes
kubectl get deployments
kubectl get services
kubectl get ingress
kubectl get endpoints
kubectl get namespaces
kubectl get hpa --all-namespaces

kubectl get pod my-pod -o yaml
kubectl get pod my-pod -o wide
kubectl get pods mypod -o jsonpath='{.spec.containers[*].name}'

kubectl get events
kubectl config view
kubectl config get-clusters
kubectl config get-contexts
kubectl config use-context CONTEXT_NAME

kubectl delete pod nginx
kubectl scale deployment hello-node --replicas 2

kubectl describe nodes node1
kubectl describe pods pod1
kubectl describe deployments deployment1 -n namespace
kubectl describe services service1
kubectl describe ingress ingres1

kubectls top nodes  # return node's cpu & mem usages(%, exact)
kubectls top pod    # return pod's exact cpu & mem usages
kubectl top pod | grep "podRegx1\|podRegx2"

kubectl logs my-pod -c my-container
kubectl exec -it mypod bash -c mycontainer

kubectl create deployment hello-node --image=registry.k8s.io/echoserver:1.4
kubectl apply -f ./nginx-deployment.yaml # yaml file can be web url
kubectl expose deployment hello-node --type=LoadBalancer --port=8080 # expose this port app is listening on

kubectl get pods --show-labels --all-namespaces  | grep ingress | grep controller
kubectl autoscale deployment php-apache --cpu-percent=50 --min=1 --max=10
```

```
kubectx
stern podRegex -c containerRegex
```
## helm & chart
* [how to use helm3 to install software on k8](https://www.digitalocean.com/community/tutorials/how-to-install-software-on-kubernetes-clusters-with-the-helm-3-package-manager)
* chart bootstrap resource i.e telegraf deployment on k8 cluster using helm pkg manager

```
#example-chart i.e ngnix
charts/
templates/
├─ tests/
│  ├─ test-connection.yaml
├─ deployment.yaml
├─ hpa.yaml
├─ ingress.yaml
├─ NOTES.txt
├─ service.yaml
├─ serviceaccount.yaml
├─ _helpers.tpl
Chart.yaml
values.yaml
```

```
#cat example-chart/templates/service.yaml
apiVersion: v1
kind: Service
metadata:
  name: {{ include "mychart.fullname" . }}
  labels:
    {{- include "mychart.labels" . | nindent 4 }}
spec:
  type: {{ .Values.service.type }}
  ports:
    - port: {{ .Values.service.port }}
      targetPort: http
      protocol: TCP
      name: http
  selector:
    {{- include "mychart.selectorLabels" . | nindent 4 }}
```

```
helm install ingress-nginx/ingress-nginx --set variable_name=variable_value
helm list
helm upgrade ingress-nginx ingress-nginx/ingress-nginx --set controller.replicaCount=3 --reuse-values

helm install example-chart --dry-run --debug ./example-chart
helm package ./example-chart
helm install example-chart example-chart-0.1.0.tgz
```

## Getting started

* tool
  * `Minikube`: very simple, minimal. single node cluster. no way to add other node.
  * `Kubeadm`: can add 1 master, N worker node. Require powerfull laptop.
  * `kind`: It support all sort of cluster. N master, N worker. require docker. complicated external networking.
  * `k3s`: lightweight version of kubernetes. some feature missing. uses sqlite as db compare to etcd/postgres in k8
  * `k9s`:
  * `rancher`: 


## Kubernetes Metrics
* 4 golden rule: `latency`, `traffic`, `errors`, `saturation`(underlying cpu, mem etc)
* `node_exporter` a prometheus client sends 1000s of metric at node level.
* `cAdvisor` embeded in kubelet, provide metric at container level
* `limit` and `request` are set at container level. If cpu exceed limit then its throttled. If memory exceed limit then it is killed. 

### node cpu utilization: 
*  `node_cpu` is counter, so rate is needed.
* `node_load1` is gauge, 1 minute avg load. Need cpu core in order to make sense.
* we can derive count of cpu: `count(node_cpu{mode="system"}) by (node)`

```
sum(node_load1) by (node) / count(node_cpu{mode="system"}) by (node) * 100
```
```
sum(rate(
         node_cpu{mode!=”idle”,
                  mode!=”iowait”,
                  mode!~”^(?:guest.*)$”
                  }[5m])) BY (instance)
```

### node memory utilization
* Apart from `free`, `buffer` and `cached` are also free memory. So 
```
sum(node_memory_MemFree + node_memory_Cached + node_memory_Buffers)
```
* Newer linux kernel reporter better free memory metric: `node_memory_MemAvailable`
```
1 - sum(node_memory_MemAvailable) by (node) 
/ sum(node_memory_MemTotal) by (node)
```

### node disk utilization
* `node_disk_io_now`, `node_disk_io_time_ms`, `node_disk_io_weighted` 
```
sum(node_filesystem_free{mountpoint="/"}) by (node, mountpoint) / sum(node_filesystem_size{mountpoint="/"}) by (node, mountpoint)
```

### node network utilization
```
sum(rate(node_network_receive_bytes[5m])) by (node) + sum(rate(node_network_transmit_bytes[5m])) by (node)
```

```
sum(rate(node_network_receive_drop[5m])) by (node) + sum(rate(node_network_transmit_drop[5m])) by (node)
```

### container cpu utilization
* `container_cpu_user_seconds_total`: total user time
* `container_cpu_system_seconds_total`: total system time
* `container_cpu_usage_seconds_total`: sum of above
* all above are counter, so apply rate.
```
sum(
    rate(container_cpu_usage_seconds_total[5m]))
by (container_name)
```

```
sum(container_memory_working_set_bytes) by (container_name) / sum(label_join(kube_pod_container_resource_limits_memory_bytes,
    "container_name", "", "container")) by (container_name)
```

## Minikube
* container runtime is prerequiste like docker, podman, hyperkit etc
* On cloud providers that support load balancers, an external IP address would be provisioned to access the Service. On minikube, the LoadBalancer type makes the Service accessible through the minikube service command.
```
brew install minikube
minikube start --driver=docker

kubectl create deployment hello-node --image=registry.k8s.io/echoserver:1.4
kubectl expose deployment hello-node --type=LoadBalancer --port=8080
minikube service hello-node

minikube addons list
minikube addons enable metrics-server
minikube addons disable metrics-server

kubectl delete service hello-node
kubectl delete deployment hello-node
minikube stop
minikube delete # delete vm, avoid this unless needed again.

minikube config set cpu <whatever>
minikube config set memory <whatever>
```

## Kind
* docker is prerequisite
```
brew install kind
kind create cluster #default one
kind create cluster --name kind-2
kind get clusters
```

## Helm
* helm chart template are written in go template language.
* values can be provided via `values.yaml` file or at command line via `helm install -f values.yaml`
* To override values in a chart, use either the '--values' flag and pass in a file or use the '--set' flag
```
brew install helm

helm repo add helm http://chartmuseum.golabs.io/stable
helm repo add bitnami https://charts.bitnami.com/bitnami

helm install jenkins ./jenkins-1.2.3.tgz
helm install jenkins-deployment ./jenkins-archive
helm install jenkins https://example.com/charts/jenkins-1.2.3.tgz
helm install --repo https://example.com/charts/ jenkins-deployment jenkins

helm search repo bitnami
helm show chart bitnami/mysql
helm show all bitnami/mysql
helm repo update
helm install bitnami/mysql --generate-name
helm uninstall mysql-1612624192
helm list                                   #list deployed release

helm upgrade --set foo=bar --set foo=newbar -f myvalues.yaml -f override.yaml redis ./redis
helm upgrade --install ${CI_PROJECT_NAME} url_dir_loc
      --set labels.application="service_name"
      --set image.bucketName="$(echo ${CI_PROJECT_NAME} | sed s/-/_/g)"
      --set image.tag="${APP_VERSION}-${CI_BUILD_REF}"
      --set application.entryPoint='{./app, server}'
      --set application.migrationCommand='{echo}'
      --set service.type="LoadBalancer"
      --set deployment.replicas="1"
      --set resources.limits.cpu="2000m"
      --set resources.limits.memory="2Gi"
      --set resources.requests.cpu="1000m"
      --set resources.requests.memory="1Gi"
      --set hpa.enabled="true"
      --set hpa.minReplica="1"
      --set hpa.maxReplica="2"
      --set hpa.targetCPUUtilizationPercentage="50"
      -f ./service_env_alues.yaml

```

### chart file structure
```
wordpress/
  Chart.yaml          # A YAML file containing information about the chart
  LICENSE             # OPTIONAL: A plain text file containing the license for the chart
  README.md           # OPTIONAL: A human-readable README file
  values.yaml         # The default configuration values for this chart
  values.schema.json  # OPTIONAL: A JSON Schema for imposing a structure on the values.yaml file
  charts/             # A directory containing any charts upon which this chart depends.
  crds/               # Custom Resource Definitions
  templates/          # A directory of templates that, when combined with values,
                      # will generate valid Kubernetes manifest files.
  templates/NOTES.txt # OPTIONAL: A plain text file containing short usage notes
```

### `Chart.yaml` file
```
apiVersion: The chart API version (required)
name: The name of the chart (required)
version: A SemVer 2 version (required)
kubeVersion: A SemVer range of compatible Kubernetes versions (optional)
description: A single-sentence description of this project (optional)
type: The type of the chart (optional)
keywords:
  - A list of keywords about this project (optional)
home: The URL of this projects home page (optional)
sources:
  - A list of URLs to source code for this project (optional)
dependencies: # A list of the chart requirements (optional)
  - name: The name of the chart (nginx)
    version: The version of the chart ("1.2.3")
    repository: (optional) The repository URL ("https://example.com/charts") or alias ("@repo-name")
    condition: (optional) A yaml path that resolves to a boolean, used for enabling/disabling charts (e.g. subchart1.enabled )
    tags: # (optional)
      - Tags can be used to group charts for enabling/disabling together
    import-values: # (optional)
      - ImportValues holds the mapping of source values to parent key to be imported. Each item can be a string or pair of child/parent sublist items.
    alias: (optional) Alias to be used for the chart. Useful when you have to add the same chart multiple times
maintainers: # (optional)
  - name: The maintainers name (required for each maintainer)
    email: The maintainers email (optional for each maintainer)
    url: A URL for the maintainer (optional for each maintainer)
icon: A URL to an SVG or PNG image to be used as an icon (optional).
appVersion: The version of the app that this contains (optional). Needn't be SemVer. Quotes recommended.
deprecated: Whether this chart is deprecated (optional, boolean)
annotations:
  example: A list of annotations keyed by name (optional).
```

### `values.yaml` file
```
imageRegistry: "quay.io/deis"
dockerTag: "latest"
pullPolicy: "Always"
storage: "s3"
```

### Template file
```
apiVersion: v1
kind: ReplicationController
metadata:
  name: deis-database
  namespace: deis
  labels:
    app.kubernetes.io/managed-by: deis
spec:
  replicas: 1
  selector:
    app.kubernetes.io/name: deis-database
  template:
    metadata:
      labels:
        app.kubernetes.io/name: deis-database
    spec:
      serviceAccount: deis-database
      containers:
        - name: deis-database
          image: {{ .Values.imageRegistry }}/postgres:{{ .Values.dockerTag }}
          imagePullPolicy: {{ .Values.pullPolicy }}
          ports:
            - containerPort: 5432
          env:
            - name: DATABASE_STORAGE
              value: {{ default "minio" .Values.storage }}
```

## Storage


## Network

## kubernetes components
![k8-components](./k8.png)

* Control plane components
    * kube-apiserver
    * ectc
    * kube-scheduler
    * kube-controller-manager
    * cloud-controller-manager
* Node components
    * kubelet: agent that run on each node in cluster. it make sure containers are running inpod.
    * kube-proxy: network proxy that run on each node in cluster,  implementing part of the Kubernetes Service concept.. Maintain network rules that allow comms to pods from inside & outside cluster.
    * container runtime: responsible for running container. it supports runtime such as containered, CRI-O, other impl of CRI


## Kubernetes GCP
Single zone clusters
* master and nodes both run in only one zone
* if cluster name does have zone suffix then it is zonal cluster. i.e `asia-east1-a`, `asia-east1-b`
* if zone outage then service outage

Multi zonal clusters
* master run  in only one zone
* nodes distributed among multiple zone.
* if cluster name does have zone suffix then it is zonal cluster. i.e `asia-east1-a`, `asia-east1-b`
* if master zone outage, service will still be working though won't be able to configure until master is up

Regional clusters
* master distributed among multiple zone
* nodes distrubuted among multiple zone
* if cluster name does not have zone suffix then it is reginal cluster. i.e `asia-east1`
* default no of zone: 3
* default no of nodes per zone: 3
* default no of masters: 3

## Kubernetes AWS