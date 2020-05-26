### info
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

### list of addon
minikube addons list  
minikube addons enable metrics-server
kubectl get pod,svc -n kube-system
minikube addons disable metrics-server