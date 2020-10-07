#!/usr/bin/env bash

go test github.com/aurumcodex/jdkenv/util && \
go test -cover github.com/aurumcodex/jdkenv/util && \
printf "\n" && \
CGO_ENABLE=0 go build -ldflags="-s -w" && \
./jdkenv help && \
printf "\n[[ compilation succeeded ]]\n\n"