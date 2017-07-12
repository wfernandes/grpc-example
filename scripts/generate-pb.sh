#!/bin/bash
set -ex
PROJECT_DIR="$( cd "$(dirname "$0")/.."; pwd )"

DIR=${PROJECT_DIR}/definitions

protoc -I ${DIR}/  ${DIR}/*.proto --go_out=plugins=grpc:definitions
