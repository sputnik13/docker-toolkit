#!/usr/bin/env bash
docker run --rm\
    -v "$PWD":/usr/src/myapp:Z \
    -w /usr/src/myapp \
    -e CGO_ENABLED=0 \
    -e GOOS=linux \
    -e GOARCH=amd64 \
    golang:1.8 \
    go build -ldflags="-s -w" uberExec.go

docker build -t uberexec .
