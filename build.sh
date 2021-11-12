#!/bin/bash

export CGO_ENABLED=0
export GO111MODULE=on
export PATH=/usr/local/go/bin:$PATH
go mod tidy
go build -o www
