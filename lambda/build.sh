#!/bin/bash
set -euxo pipefail
GOOS=linux go build -o hello . && zip lambda.zip hello
