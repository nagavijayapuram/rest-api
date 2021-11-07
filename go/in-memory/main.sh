#!/usr/bin/env bash

# --------
# main.sh
# --------

cd `dirname $0`

trap "/bin/rm -f main" EXIT SIGINT SIGTERM

nc -zv localhost 8000 2>/dev/null

if [ $? -eq 0 ]; then
  echo -e "\nmain is already up and running with pid `ps -ef | awk '/main$/ {print $2}'`\n"
  exit 1
fi

go build main.go && ./main &

sleep 1

echo -e "\n . Getting posts ...\n"

curl -s localhost:8000/posts | jq

echo
