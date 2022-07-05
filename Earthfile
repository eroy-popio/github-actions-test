VERSION 0.6
FROM golang:1.15-alpine3.13
WORKDIR /go-example

build:
    COPY main.go .
    RUN go build -o build/go-example main.go
    SAVE ARTIFACT build/go-example /go-example

docker:
    COPY +build/go-example .
    ENTRYPOINT ["/go-example/go-example"]
    SAVE IMAGE go-example:latest
