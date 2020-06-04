### reinstall minikube
cd ~  
minikube delete
minikube stop  
rm -rf .minikube  

minikube kubectl -- get pods
minikube version
minikube version --client
minikube get componentstatus //- показать состояние к8s кластера

minikube kubectl -- get pods // - показать все серверы кластера

### start
minikube start  
minikube -p mycluster  
minikube start --cpus=1 --memory=1gb --disk-size=2gb
minikube start --vm-driver=docker --alsologtostderr

### dashboard
minikube dashboard

### ssh
minikube ssh
user: root
or user: docker; pass: tcuser

#### configs
~/.minikube
~/.kubectl



