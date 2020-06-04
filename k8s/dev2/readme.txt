docker build -t go-docker-image .
docker run -p 8081:8081 go-docker-image

docker tag go-docker-image:latest jedi108/gogo:go-docker-image
docker push jedi108/gogo:go-docker-image
kubectl run hello --generator=run-pod/v1 --image=jedi108/gogo:go-docker-image --port=8081
kubectl port-forward hello 8083:8081

kubectl run pgsql-postgresql-client --rm --tty -i --restart='Never' --namespace default --image docker.io/bitnami/postgresql:11.7.0-debian-10-r9 --env="PGPASSWORD=$POSTGRES_PASSWORD" --command -- psql testdb --host demo-postgresql -U postgres -d postgres -p 5432
kubectl run pgsql-postgresql-client --rm --tty -i --restart='Never' --namespace default --image docker.io/bitnami/postgresql:11.7.0-debian-10-r9 --env="PGPASSWORD=PG" --env="PGUSER=PG" --command -- psql testdb --host demo-postgresql -U postgres -d postgres -p 5432

kubectl run postgresql-client -generator=run-pod/v1 --image docker.io/bitnami/postgresql:11.7.0-debian-10-r9 --env="PGPASSWORD=PG" --command -- psql testdb -U postgres -d postgres -p 5432
