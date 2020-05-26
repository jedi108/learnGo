### reinstall minikube
cd ~  
minikube delete
minikube stop  
rm -rf .minikube  
minikube start  

minikube start --cpu=1 --memory=1gb --disk-size=2gb
