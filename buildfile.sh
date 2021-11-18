#!/bin/bash

go build chatapi.go
echo "\n chatapi.go build "

pids=$(ps -eo comm,pid | awk '/^chatapi/' | awk '{print $2}')

if [ -z "$pids" ]
then
    echo "empty"
else
    kill -9 $pids
    echo kill -9 $pids
fi

