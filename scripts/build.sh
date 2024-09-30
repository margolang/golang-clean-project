#!/usr/bin/env bash
CGO_ENABLED=0 GOOS=$1 GOARCH=amd64 go build -tags=jsoniter -gcflags "all=-N -l" -o ./bin/server ./cmd/server
