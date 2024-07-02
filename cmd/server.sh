#!/bin/bash

# kill the server only if web flag is passed
if [ "$1" == "-web" ]; then
    kill-port 8080
    go run . -web
else
    go run . "$@"
fi
