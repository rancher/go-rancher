#!/bin/bash
set -e

cd $(dirname $0)/../generator

URL_BASE='http://localhost:8080'

if [ "$1" != "" ]; then
    URL_BASE=$1
fi

echo -n Waiting for cattle ${URL_BASE}/ping
while ! curl -fs ${URL_BASE}/ping; do
    echo -n .
    sleep 1
done
echo

curl -s "${URL_BASE}/v2-beta/schemas?_role=service" | jq . > schemas.json
echo Saved schemas.json

echo -n Generating go code...
go run generator.go
echo " Done"

gofmt -w ../client/generated_*
echo Formatted code

echo Success
