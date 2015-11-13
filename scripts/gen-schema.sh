#!/bin/bash
set -e

cd $(dirname $0)/../generator

curl -s 'http://localhost:8080/v1/schemas?_role=service' | jq . > schemas.json
echo Saved schemas.json

echo -n Generating go code...
godep go run generator.go
echo " Done"

gofmt -w ../client/generated_*
echo Formatted code

echo Success
