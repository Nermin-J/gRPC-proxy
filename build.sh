#!/bin/sh

ROOT=grpc-app
PROTO_DIR=$ROOT/protos
PROTOC=/usr/bin/protoc 
GO_OUT=$ROOT/stubs

rm -rf $GO_OUT
mkdir -p $GO_OUT

protoc -I$PROTO_DIR --go_out=$GO_OUT --go_opt=paths=source_relative --go-grpc_out=$GO_OUT --go-grpc_opt=paths=source_relative $(find $PROTO_DIR -type f -name '*.proto')