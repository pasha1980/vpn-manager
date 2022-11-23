#!/bin/bash
docker stop test || true
docker rm test || true

docker build -f devops/docker/Dockerfile -t test .

docker volume create openvpn_data
docker volume create openvpn_config
docker run --cap-add=NET_ADMIN \
-v openvpn_data:/opt/openvpn_data \
-v openvpn_config:/etc/openvpn \
-p 1194:1194/udp \
-p 8080:80/tcp \
--rm \
--name test \
test "$@"