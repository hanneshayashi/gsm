#!/bin/bash
APP_VERSION="0.2.6-pre"
ARCHIVE_NAME=${APP_NAME}-${APP_VERSION}
VERSIONS=("win-amd64" "linux-amd64" "linux-arm64" "mac-amd64" "mac-arm64")

mkdir tmp
for i in ${!VERSIONS[@]}; do
    cp ./LICENSE ${APP_NAME}/${VERSIONS[$i]}/
    tar -czf tmp/${ARCHIVE_NAME}_${VERSIONS[$i]}.tar.gz -C ${APP_NAME}/${VERSIONS[$i]} .
done

gsutil -m cp tmp/* gs://${BUCKET}/${APP_NAME}/

for i in ${!VERSIONS[@]}; do
    gsutil setmeta -h "Content-Disposition:filename=${ARCHIVE_NAME}_${VERSIONS[$i]}.tar.gz" gs://${BUCKET}/${APP_NAME}/${ARCHIVE_NAME}_${VERSIONS[$i]}.tar.gz
done