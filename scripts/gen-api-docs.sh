#!/bin/bash
set -e

cd $(dirname $0)/../docs

curl -s http://localhost:8080/v1/schemas?_role=project | jq . > ./input/schemas.json
echo Saved schemas.json

godep go run *.go -command=generate-description

godep go run *.go -command=generate-docs

echo Success
