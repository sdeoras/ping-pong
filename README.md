## About
`ping-pong` is a `gRPC` based client/server implimentation to test micro-service
communication within your distributed computing environment. Client sends `ping` to the server and the server responds with a `pong` and the game continues.

If you are using `skaffold` you will be able to run it directly, otherwise, you will need to modify image names in the kubernetes manifests.

Check client and server logs.

## Prerequisites
This repo requires `grpc` setup. From your home folder:
* [Install protoc](https://developers.google.com/protocol-buffers/docs/downloads)

Install go libs for grpc and go plugin
```bash
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Setup
This repo used `go modules`. Clone this repo in a location outside `GOPATH`
```bash
git clone https://github.com/sdeoras/ping-pong.git
```

`go get` dependencies
```bash
go get ./...
```

## Known Issues
`skaffold` does not work properly with `cri-o` runtime managed k8s clusters. Pl. see this
[issue](https://github.com/cri-o/cri-o/issues/2351). Try following yaml to see if the issue
got resolved.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: busybox-no-sha
  labels:
    name: busybox-no-sha
spec:
  restartPolicy: OnFailure
  containers:
  - name: busybox-no-sha
    image: busybox:latest
    imagePullPolicy: Always
    command:
      - echo
      - no-sha
    resources:
      limits:
        memory: "128Mi"
        cpu: "500m"
    ports:
      - containerPort: 8080
---
apiVersion: v1
kind: Pod
metadata:
  name: busybox-sha
  labels:
    name: busybox-sha
spec:
  restartPolicy: OnFailure
  containers:
  - name: busybox-sha
    image: busybox:latest@sha256:4b6ad3a68d34da29bf7c8ccb5d355ba8b4babcad1f99798204e7abb43e54ee3d
    imagePullPolicy: Always
    command:
      - echo
      - sha
    resources:
      limits:
        memory: "128Mi"
        cpu: "500m"
    ports:
      - containerPort: 8080
```