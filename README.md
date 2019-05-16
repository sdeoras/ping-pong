## A Simple client-server ping-pong to test
This repo defines a `client` that sends `ping` to a `server` and the `server`
responds with a `pong`. This was written to quickly test micro-services
communication in your distributed computing environment.

## Prerequisites
This repo requires `grpc` setup.

* [Install protoc](https://developers.google.com/protocol-buffers/docs/downloads))

Install go libs for grpc and go plugin
```bash
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Setup
Clone this repo in a location outside `GOPATH`
```bash
git clone https://github.com/sdeoras/ping-pong.git
```

`go get` dependencies
```bash
go get ./...
```
