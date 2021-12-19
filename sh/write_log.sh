#!/bin/bash

tf="/tmp/tailpipe_test.log"

arr=(
    "debug" "info" "warn" "error" "fatal"
)

echo "Write test messages to ${tf}"
# test script generating a log file
while true; do
    r=$((RANDOM % 5 + 1))
    sleep ${r}
    echo -e "[${arr[$r - 1]}]\t$(date)" | tee -a "${tf}"
done
