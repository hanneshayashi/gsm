#!/bin/bash
cp ./LICENSE gsm/win-x64/
cp ./LICENSE gsm/linux-x64/
cp ./LICENSE gsm/linux-arm64/
cp ./LICENSE gsm/mac-x64/

mkdir tmp

tar -czf tmp/gsm_win-x64.tar.gz -C gsm/win-x64 .
tar -czf tmp/gsm_linux-x64.tar.gz -C gsm/linux-x64 .
tar -czf tmp/gsm_linux-arm64.tar.gz -C gsm/linux-arm64 .
tar -czf tmp/gsm_mac-x64.tar.gz -C gsm/mac-x64 .

gsutil -m cp tmp/* gs://build-arts/gsm/

gsutil setmeta -h 'Content-Disposition:filename=gsm_win-x64.tar.gz' gs://build-arts/gsm/gsm_win-x64.tar.gz
gsutil setmeta -h 'Content-Disposition:filename=gsm_linux-x64.tar.gz' gs://build-arts/gsm/gsm_linux-x64.tar.gz
gsutil setmeta -h 'Content-Disposition:filename=gsm_linux-arm64.tar.gz' gs://build-arts/gsm/gsm_linux-arm64.tar.gz
gsutil setmeta -h 'Content-Disposition:filename=gsm_mac-x64.tar.gz' gs://build-arts/gsm/gsm_mac-x64.tar.gz

# gsutil -m acl ch -r -g AllUsers:R gs://build-arts/gsm/*