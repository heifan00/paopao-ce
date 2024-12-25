#!/bin/sh
# eg.1 : sh build-image.sh
# eg.2, set tags: sh build-image.sh 'go_json'

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o release/linux-amd64/paopao

# nohup ./paopao serve > paopao.log 2>&1 &
