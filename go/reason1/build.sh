#!/bin/bash
go build -ldflags="-X 'main.environment=production'" main.go

docker build --build-arg ENV=production -t main .
