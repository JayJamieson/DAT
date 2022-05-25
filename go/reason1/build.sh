#!/bin/bash
go build -ldflags="-X 'main.environment=production'" main.go
