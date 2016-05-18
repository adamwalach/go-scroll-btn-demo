#!/bin/bash

go version
GOOS=linux GOARCH=arm GOARM=5 CGO_ENABLED=0 go build -o=main *.go
