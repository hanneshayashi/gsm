#!/bin/bash
cp ./LICENSE gsm/win-amd64/
cp ./LICENSE gsm/linux-amd64/
cp ./LICENSE gsm/linux-arm64/
cp ./LICENSE gsm/mac-amd64/
cp ./LICENSE gsm/mac-arm64/

mkdir tmp

tar -czf tmp/gsm_win-amd64.tar.gz -C gsm/win-amd64 .
tar -czf tmp/gsm_linux-amd64.tar.gz -C gsm/linux-amd64 .
tar -czf tmp/gsm_linux-arm64.tar.gz -C gsm/linux-arm64 .
tar -czf tmp/gsm_mac-amd64.tar.gz -C gsm/mac-amd64 .
tar -czf tmp/gsm_mac-arm64.tar.gz -C gsm/mac-arm64 .

gsutil -m cp tmp/* gs://build-arts/gsm/

gsutil setmeta -h 'Content-Disposition:filename=gsm_win-amd64.tar.gz' gs://build-arts/gsm/gsm_win-amd64.tar.gz
gsutil setmeta -h 'Content-Disposition:filename=gsm_linux-amd64.tar.gz' gs://build-arts/gsm/gsm_linux-amd64.tar.gz
gsutil setmeta -h 'Content-Disposition:filename=gsm_linux-arm64.tar.gz' gs://build-arts/gsm/gsm_linux-arm64.tar.gz
gsutil setmeta -h 'Content-Disposition:filename=gsm_mac-amd64.tar.gz' gs://build-arts/gsm/gsm_mac-amd64.tar.gz
gsutil setmeta -h 'Content-Disposition:filename=gsm_mac-arm64.tar.gz' gs://build-arts/gsm/gsm_mac-arm64.tar.gz