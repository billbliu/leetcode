#!/usr/bin/env bash

set -e
echo "" > coverage.txt

for d in $(go list ./problems/... | grep -v vendor); do
    echo $d
    go test -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done