#!/usr/bin/env bash

if ! command -v runsc &> /dev/null
then
    echo "runsc could not be found"
    exit
fi

docker build -t skc-sandbox ./execution
