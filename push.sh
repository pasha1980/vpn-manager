#!/bin/bash
docker build -f devops/docker/Dockerfile -t registry.gitlab.com/khvalygin/tgvpn/dev:latest .
docker push registry.gitlab.com/khvalygin/tgvpn/dev:latest