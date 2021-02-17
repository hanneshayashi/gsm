#!/bin/bash
mkdir gsm
mkdir gsm/win-amd64
mkdir gsm/linux-amd64
mkdir gsm/linux-arm64
mkdir gsm/mac-amd64
mkdir gsm/mac-arm64

go version

GOARCH=amd64 GOOS=windows go build -o gsm/win-amd64/gsm.exe
GOARCH=amd64 GOOS=linux go build -o gsm/linux-amd64/gsm
GOARCH=arm64 GOOS=linux go build -o gsm/linux-arm64/gsm
GOARCH=amd64 GOOS=darwin go build -o gsm/mac-amd64/gsm
GOARCH=arm64 GOOS=darwin go build -o gsm/mac-arm64/gsm