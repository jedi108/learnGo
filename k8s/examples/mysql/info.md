https://kubernetes.io/docs/tutorials/stateful-application/mysql-wordpress-persistent-volume/

## deploy
kubectl get deployments
#### mysql
kubectl apply -f examples/mysql/mysql_pv.yaml   
kubectl apply -f examples/mysql/mysql_deployment.yaml

### run
kubectl run -it --rm --image=mysql:5.6 --restart=Never mysql-client -- mysql -h mysql -ppassword

### serrets
kubectl apply -k examples/mysql/

### delete all
kubectl delete -k ./  
