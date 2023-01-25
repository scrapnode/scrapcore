#!/bin/bash

set -e

export CI=true

mkdir -p coverage
go test -tags test -v -coverprofile=coverage/cover.out.tmp ./...
# shellcheck disable=SC2002
cat coverage/cover.out.tmp | grep -v "tester.go" > coverage/cover.out
go tool cover -func=coverage/cover.out -o coverage/coverage.txt

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
bash "${__dir}/ci_validate_coverage.sh"