#!/usr/bin/env bash

# --------
# main.sh
# --------

go run main.go 2>/dev/null &

sleep 1

echo -e "\n . Getting posts ...\n"

curl -s localhost:8000/posts | jq

echo
