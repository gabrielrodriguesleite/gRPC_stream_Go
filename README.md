# gRPC_stream_Go
A repo to train gRPC stream with Golang

Some stream can be handfull to help to send a big amount of data in tiny chunks.

## Config a Go module and populate a gRPC simple project

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

## Create a .proto definition

After define a .proto execute

`make generate`

To build `data.grpc.pg.go` and `data.pg.go` packages

###### IF IT THROWS A ERROR

Includes this line on the very end of file `~/.bashrc` or `~/.zshrc`

```sh
# PROTO GO plugins
export PATH=$PATH:$(go env GOPATH)/bin
```

Reload environment variables with `source ~/.zhrc/` or relogin on the session.

## Testing

## Go fix dependencies

To get all the dependencies for grpc proto run:

`go mod tidy`

A `go.sum` file should be automaticaly created

###### REFERENCES
https://jadekler.github.io/2020/10/08/stubbing-grpc.html