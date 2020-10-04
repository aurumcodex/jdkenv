#!/usr/bin/env bash

go version
CGO_ENABLE=0 go build -ldflags="-s -w"
./jdkenv help