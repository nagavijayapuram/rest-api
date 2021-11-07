#!/usr/bin/env bash

# --------
# main.sh
# --------

cd `dirname $0`

trap "/bin/rm -f main" EXIT SIGINT SIGTERM

nc -zv localhost 8000 2>/dev/null

if [ $? -eq 0 ]; then
  pid=`ps -ef | awk '/var.*mai[n]/ {print $2}'`
  echo -en "\n . Killing the process from previous run - "
  kill -9 $pid
fi

go run main.go &

sleep 1

echo -e "\n . Getting posts ...\n"

curl -s localhost:8000/posts | jq

echo
