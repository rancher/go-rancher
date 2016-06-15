#!/bin/bash
set -e

CATTLE_VERSION=v0.161.3

cleanup()
{
    e=$?
    if [ -n "$PID" ]; then
        kill $PID
    fi
    exit $e
}

trap cleanup EXIT

cd $(dirname $0)/../generator

CATTLE_JAR=https://github.com/rancher/cattle/releases/download/${CATTLE_VERSION}/cattle.jar
URL_BASE='http://localhost:8080'

if [ "$1" != "-l" ]; then
    URL_BASE='http://localhost:18282'
    docker run -p 18282:8080 --rm -e CATTLE_MACHINE_EXECUTE=false -e URL=$CATTLE_JAR rancher/server:v1.0.1-rc1 &
    PID=$!
    sleep 2
fi

echo -n Waiting for cattle ${URL_BASE}/ping
while ! curl -fs ${URL_BASE}/ping; do
    echo -n .
    sleep 1
done
echo

curl -s "${URL_BASE}/v1/schemas?_role=service" | jq . > schemas.json
echo Saved schemas.json

echo -n Generating go code...
godep go run generator.go
echo " Done"

gofmt -w ../client/generated_*
echo Formatted code

echo Success
