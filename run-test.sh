#!/bin/bash
docker stop test || true
docker rm test || true

docker build -f devops/docker/Dockerfile -t test .
docker run --cap-add=NET_ADMIN \
-v ./tmp/openvpn_data:/opt/openvpn_data \
-v ./tmp/openvpn_config:/etc/openvpn \
-p 1194:1194/udp \
-p 8080:80/tcp \
-e HOST_ADDR=127.0.0.1 \
-e HOST_URL=127.0.0.1:8080 \
--rm \
--name test \
test "$@"