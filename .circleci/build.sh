#!/bin/bash

docker build -t farhansajid2/user-service:$CIRCLE_SHA1 ./user-service/Dockerfile .
docker build -t farhansajid2/email-service:$CIRCLE_SHA1 ./email-service/Dockerfile .

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_ID" --password-stdin

docker push farhansajid2/email-service:$CIRCLE_SHA1
docker push farhansajid2/user-service:$CIRCLE_SHA1