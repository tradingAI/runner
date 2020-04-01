#!/bin/sh

set -e

ROOT="$GOPATH/src/github.com/tradingAI"
PROJECT_ROOT="$ROOT/runner"
PROTO_ROOT="$ROOT/proto"
PROTO_GEN_DIR="$PROTO_ROOT/gen"

# clear old
rm -rf $PROTO_GEN_DIR

for element in `ls $PROTO_ROOT`
  do
      if [ -d $PROTO_ROOT/$element ];then
          protoc \
              -I $PROTO_ROOT \
              --go_out=plugins=grpc:$GOPATH/src \
              $PROTO_ROOT/$element/*.proto
      fi
  done
