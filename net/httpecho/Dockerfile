FROM alpine:3.7
ARG COMPONENT_VERSION=undefined

ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0
#ENV GONOPROXY stash.lamoda.ru
#ENV GONOSUMDB stash.lamoda.ru
#ENV GOPRIVATE stash.lamoda.ru
ENV VERSION ${COMPONENT_VERSION}

WORKDIR /go/src/github.com/jedi108/learnGo/

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

# This is the ‘magic’ step that will download all the dependencies that are specified in
# the go.mod and go.sum file.
# Because of how the layer caching system works in Docker, the go mod download
# command will _only_ be re-run when the go.mod or go.sum file change
RUN go mod download
COPY . .

RUN make build VERSION=${VERSION}

FROM alpine:3.9
RUN apk add --no-cache --virtual tzdata ca-certificates
COPY --from=build /go/src/github.com/jedi108/learnGo /bin/echoserver
ENTRYPOINT ["/bin/echoserver"]