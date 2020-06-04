### docker 
docker build -t simplehttp net/httpecho/

go build -o simplehttp github.com/jedi108/learnGo/net/httpecho/

docker image rm simplehttp -f

docker run -p 8081:8081 --name simplehttp --rm simplehttp