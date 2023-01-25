#!/bin/bash

set -e

go test -tags test -v -coverprofile=coverage/cover.out.tmp ./...
go tool cover -html coverage/cover.out.tmp -o coverage/cover.out.html
