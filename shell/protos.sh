#!/bin/bash

# 获取脚本所在的目录
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# 如果脚本在外部文件夹中执行，设置输出路径为./pkg
if [ "$SCRIPT_DIR" != "$(pwd)" ]; then
  OUTPUT_DIR="./pkg"
  PROTO_PATH="./protos/"
  PROTO_PATH_FILE="./protos/*.proto"
else
  OUTPUT_DIR="../pkg"
  PROTO_PATH="../protos/"
  PROTO_PATH_FILE="../protos/*.proto"
fi

# 执行 protoc 命令
protoc --go_out=plugins=grpc:"$OUTPUT_DIR" --proto_path="${GOPATH}/src" --proto_path="$PROTO_PATH" $PROTO_PATH_FILE