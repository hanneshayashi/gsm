#!/bin/bash
mkdir gsm
mkdir gsm/win-x64
mkdir gsm/linux-x64
mkdir gsm/linux-arm64
mkdir gsm/mac-x64

GOARCH=amd64 GOOS=windows go build -o gsm/win-x64/gsm.exe
GOARCH=amd64 GOOS=linux go build -o gsm/linux-x64/gsm
GOARCH=arm64 GOOS=linux go build -o gsm/linux-arm64/gsm
GOARCH=amd64 GOOS=darwin go build -o gsm/mac-x64/gsm