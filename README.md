# gRPC_stream_Go
A repo to train gRPC stream with Golang

Some stream can be handfull to help to send a big amount of data in tiny chunks.

## Configa a Go module

```sh
# init go project
PACKAGE_NAME=grpc_stream_go
go mod init $PACKAGE_NAME

# create the folders
mkdir server client proto

# populate them
echo package main > client/main.go
echo package main > server/main.go
echo 'syntax = "proto3";' > proto/data.proto
```