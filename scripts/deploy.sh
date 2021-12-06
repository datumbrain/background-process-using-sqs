#!/usr/bin/env bash

export GOOS=linux
export GOARCH=amd64

go build -o bin/main-program .
go build -o bin/sqs-handler ./sqs

serverless deploy
