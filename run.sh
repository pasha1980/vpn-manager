#!/bin/bash
docker stop test || true
docker rm test || true

docker build -f devops/dev/Dockerfile -t test .
docker run --cap-add=NET_ADMIN \
-v data:/opt \
-p 1194:1194/udp \
-e HOST_ADDR=31.144.22.123 \
--rm \
--name test \
test "$@"