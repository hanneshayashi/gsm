#!/bin/bash
APP_VERSION="0.2.6"
BINARY_NAME=${APP_NAME}_v${APP_VERSION}
VERSIONS=("win-amd64" "linux-amd64" "linux-arm64" "mac-amd64" "mac-arm64")

mkdir ${APP_NAME}

for i in ${!VERSIONS[@]}; do
    mkdir -p ${APP_NAME}/${VERSIONS[$i]}
done

go version
GOARCH=amd64 GOOS=windows go build -o ${APP_NAME}/win-amd64/${BINARY_NAME}.exe
GOARCH=amd64 GOOS=linux go build -o ${APP_NAME}/linux-amd64/${BINARY_NAME}
GOARCH=arm64 GOOS=linux go build -o ${APP_NAME}/linux-arm64/${BINARY_NAME}
GOARCH=amd64 GOOS=darwin go build -o ${APP_NAME}/mac-amd64/${BINARY_NAME}
GOARCH=arm64 GOOS=darwin go build -o ${APP_NAME}/mac-arm64/${BINARY_NAME}