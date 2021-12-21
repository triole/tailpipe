#!/bin/bash

tf="/tmp/tailpipe_test.log"

json='{"level":"error","date":"now","msg":"something bad happened","info":"some more information"}'
text="an error message"

arr=(
    "debug" "info" "warn" "error" "fatal"
)

echo "Write test messages to ${tf}"
# test script generating a log file
while true; do
    r=$((RANDOM % 5 + 1))
    if [[ "${1}" != "--flood" ]]; then
        sleep ${r}
    fi

    msg="[${arr[$r - 1]}]\t$(date) ${text}"
    q=$((RANDOM % 3 + 1))
    if (("${q}" == 2)); then
        msg="${json}"
    fi
    echo -e "${msg}" | tee -a "${tf}"
done
