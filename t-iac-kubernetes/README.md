
## service
* it is used to expose the service
* type
    * clusterIP: reachable within the cluster
    * NodePort: extension of clusterIP. reachable from outside cluster and external traffic
    * LoadBalancer: extension of Nodeport. expose service externaly using cloud own(aws, gcp) load balancer
    * ExternalName
![](./service.png)

## Ingress
* ingress can also be used to expose service under same ip address using different name. mainly used for http service.
* it does not expose arbitrary ports or protocol. To expose services than http/https(i.e may be grpc), we require servcie `NodePort` or `LoadBalancer`.

Ingress controller: 
* it can be software load balancer or hardware or cloud load balancer running externally.
* examples: Nginx, Ambassador, EnRoute, HAProxy, AWS ALB, AKS Application Gateway

![](./ingress.png)

## commands
```
kubectl get pods
kubectl get nodes
kubectl get deployments
kubectl get services
kubectl get ingress
kubectl get pod my-pod -o yaml
kubectl get pod my-pod -o wide
kubectl get pods mypod -o jsonpath='{.spec.containers[*].name}'



kubectl describe nodes node1
kubectl describe pods pod1
kubectl describe deployments deployment1
kubectl describe services service1
kubectl describe ingress ingres1

kubectls top nodes  # return node's cpu & mem usages(%, exact)
kubectls top pod    # return pod's exact cpu & mem usages
kubectl top pod | grep "podRegx1\|podRegx2"

kubectl logs my-pod -c my-container
kubectl exec -it mypod bash -c mycontainer
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