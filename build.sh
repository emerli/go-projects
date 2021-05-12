#!/bin/bash
echo "Building project"
export GOOS=linux
export GOARCH=amd64

# compila senza C go (adapter)
export CGO_ENABLED=0

go get -u
go test
go build -a -installsuffix cgo -o build/main .
