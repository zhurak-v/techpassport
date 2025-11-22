#!/bin/bash

export DOCKER_BUILDKIT=1

docker-compose -f docker/docker-compose.yml up --build -d

#docker build -t auth-service:latest -f ./docker/Development.Dockerfile .
#
#if [ "$(docker ps -aq -f name=auth-service)" ]; then
#    docker rm -f auth-service
#fi
#
#docker run -d --name auth-service -p 1111:1111 auth-service