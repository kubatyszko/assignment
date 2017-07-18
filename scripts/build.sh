#!/bin/bash

export CGO_ENABLED=0 
export GOOS=linux 
go get -u github.com/astaxie/beego
go build -a -o assignment assignment.go

