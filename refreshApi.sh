#!/bin/bash

cd $(dirname $0)/proto/api
work_path=$(pwd)

protoc -I. -I../googleapis --go_out=plugins=grpc:../../api *.proto
protoc -I. -I../googleapis --grpc-gateway_out=logtostderr=true:../../api *.proto

echo 'Refresh Success!'