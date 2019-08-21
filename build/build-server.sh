#!/usr/bin/env bash

# environment variables
# system
export OS="linux"
export ARCH="amd64"
export COMPONENT="server"
export CONTAINER_HOSTNAME="ping-pong-${COMPONENT}"
# config
export DEFAULT_IMAGE_BASENAME="docker.io/sdeoras/${CONTAINER_HOSTNAME}"

# create a new tmp container from centos base and capture id in CTR.
CTR=$(buildah from docker.io/library/alpine:latest)
TAG=`git rev-parse HEAD | cut -c 1-8`
STATE=`[[ -n $(git status -s) ]] || echo -n 'clean'`
if [[ "${STATE}" = "" ]]; then
    STATE="dirty"
fi
IMG="${DEFAULT_IMAGE_BASENAME}:${TAG}-${STATE}"
echo "building container image: ${IMG}"

buildah config --os="linux" --arch="amd64" ${CTR}
buildah config --created-by "buildah" ${CTR}
# currently unsupported commands, but should be added in future
# buildah config --hostname ${CONTAINER_HOSTNAME}

# configure environment variables for the container runtime.
buildah config -e OS=${OS} ${CTR}
buildah config -e ARCH=${ARCH} ${CTR}

go build -o ping-pong-${COMPONENT} github.com/sdeoras/ping-pong/${COMPONENT}/cmd/src/ping-pong-${COMPONENT}
buildah copy ${CTR} ping-pong-${COMPONENT} /
rm -rf ping-pong-${COMPONENT}

buildah run ${CTR} -- mkdir /lib64 
buildah run ${CTR} -- ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

# update yaml file
mkdir -p ../kubernetes-manifests
sed -e "s/image:.*/image: docker.io\/sdeoras\/ping-pong-${COMPONENT}:${TAG}-${STATE}/g" ${COMPONENT}.yaml > ../kubernetes-manifests/${COMPONENT}.yaml

# commit working container to a container image and remove working container
buildah commit --rm ${CTR} ${IMG}

echo "====================================="
echo "built container image ${IMG}"
echo "====================================="

