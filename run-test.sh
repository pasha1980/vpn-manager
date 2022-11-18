#!/bin/bash
docker stop test || true
docker rm test || true

docker build -f devops/docker/Dockerfile -t test .
docker run --cap-add=NET_ADMIN \
-v data:/opt \
-p 1194:1194/udp \
-p 8080:80 \
-e HOST_ADDR=127.0.0.1 \
--rm \
--name test \
test "$@"