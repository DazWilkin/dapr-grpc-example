#!/usr/bin/env bash

REGISTRY="ghcr.io"
USER="dazwilkin"
REPO="dapr-grpc"
TAG="v0.0.2"

for COMPONENT in "client" "server"
do
    IMAGE=${REGISTRY}/${USER}/${REPO}-${COMPONENT}
    docker build \
    --tag=${IMAGE}:${TAG} \
    --file=./deployment/Dockerfile.${COMPONENT} \
    .
    docker push ${IMAGE}:${TAG}
done
