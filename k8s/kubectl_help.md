### info
https://kubernetes.io/ru/docs/reference/kubectl/cheatsheet/

kubectl describe pvc mysql-pv-claim  
kubectl get events  
kubectl config view   
kubectl get pod,svc -n kube-system  
kubectl get componentstatus
kubectl cluster-info

### set namespace
kubectl config set-context --current --namespace=default

### problems urls
https://issue.life/questions/52916548  

# pods, deployments, service
kubectl get deployments
kubectl get pods
kubectl get pods -l app=mysql
kubectl get pods -o wide --all-namespaces  
kubectl delete service hello-node 
kubectl delete deployment hello-node 

### network
kubectl --namespace default port-forward $POD_NAME 8088:80
kubectl explain pods

### list of addon
minikube addons list  
minikube addons enable metrics-server
kubectl get pod,svc -n kube-system
minikube addons disable metrics-server

### rename nodes
kubectl label node node2 node.role.kubernetes.io/worker= //- label: worker

### logs
Get the pod status, Command - 
kubectl get pods

Describe pod to have further look - 
kubectl describe pod "pod-name" The last few lines of output gives you events and where your deployment failed

Get logs for more details - 
kubectl logs "pod-name"

Get container logs - 
kubectl logs "pod-name" -c "container-name" Get the container name from the output of describe pod command
If your container is up, you can use the 

kubectl exec -it command
to further analyse the container

####
kubectl get cm coredns -n kube-system -o yaml
kubectl exec -it nsqlookup-nsqlookupd-0 sh
hostname -f
