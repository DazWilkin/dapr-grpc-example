#!/usr/bin/env bash

REGISTRY="ghcr.io"
USER="DazWilkin"
REPO="dapr-grpc-example"
TAG=$(git rev-parse HEAD)

for COMPONENT in "client" "server"
do
    IMAGE=${REGISTRY}/${USER}/${REPO}
    docker build \
    --tag=${IMAGE}:${TAG} \
    --file=./deployment/Dockerfile.${COMPONENT} \
    .
    docker push ${IMAGE}:${TAG}
done
