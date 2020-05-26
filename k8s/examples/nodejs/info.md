# https://kubernetes.io/ru/docs/tutorials/hello-minikube/

### deployment
kubectl create deployment hello-node --image=gcr.io/hello-minikube-zero-install/hello-node


### expose
kubectl expose deployment hello-node --type=LoadBalancer --port=8080

### publish service
minikube service hello-node  