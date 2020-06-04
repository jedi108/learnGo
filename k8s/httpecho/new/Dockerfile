FROM golang:latest 
WORKDIR /app

#RUN mkdir /app
#ADD . /app
COPY ./ /app/.
RUN ls -la && pwd
#RUN go mod download
RUN go build -o echoserver .
EXPOSE 8081
RUN ["/app/echoserver"]
#CMD ["/app/echoserver"]